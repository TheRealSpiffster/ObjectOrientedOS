package emulator

import (
	"bytes"
	"encoding/binary"
)

type Register []byte

type Segment struct {
	offset, length uint64
}

type Core struct {
	generalPurpose [16]Register
	status         StatusRegister
	// 8 segment registers
	CodeSegment      Segment
	GlobalSegment    Segment
	DataSegment      Segment
	StackSegment     Segment
	UserStackSegment Segment
	InputSegment     Segment
	OutputSegment    Segment
	SharedSegment    Segment

	codePointer uint64
	startBlock  uint64

	running   bool
	memory    *[]byte
	codeCache []byte

	DoOperation []func([]byte)
}

type Emulator struct {
	memory []byte
	cores  []Core
}

func (e *Emulator) Init() {
	// allocate 100K for memory
	e.memory = make([]byte, 102400)
	// allocate 8 cores
	e.cores = make([]Core, 8)
}

func (c *Core) Start(memory *[]byte, memLoc int) {
	c.running = true
	c.memory = memory

	for i, r := range c.generalPurpose {
		r = (*memory)[memLoc+32*i : memLoc+32*(i+1)]
		c.generalPurpose[i] = r
	}
	memLoc += 32 * 16
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, (*memory)[memLoc:memLoc+32])
	if err != nil {
		// error in write
	}
	err = binary.Read(buf, binary.LittleEndian, &c.status)
	memLoc += 32
	c.CodeSegment = Segment{binary.LittleEndian.Uint64((*memory)[memLoc : memLoc+8]), binary.LittleEndian.Uint64((*memory)[memLoc+8 : memLoc+16])}
	memLoc += 16
	c.GlobalSegment = Segment{binary.LittleEndian.Uint64((*memory)[memLoc : memLoc+8]), binary.LittleEndian.Uint64((*memory)[memLoc+8 : memLoc+16])}
	memLoc += 16
	c.DataSegment = Segment{binary.LittleEndian.Uint64((*memory)[memLoc : memLoc+8]), binary.LittleEndian.Uint64((*memory)[memLoc+8 : memLoc+16])}
	memLoc += 16
	c.StackSegment = Segment{binary.LittleEndian.Uint64((*memory)[memLoc : memLoc+8]), binary.LittleEndian.Uint64((*memory)[memLoc+8 : memLoc+16])}
	memLoc += 16
	c.UserStackSegment = Segment{binary.LittleEndian.Uint64((*memory)[memLoc : memLoc+8]), binary.LittleEndian.Uint64((*memory)[memLoc+8 : memLoc+16])}
	memLoc += 16
	c.InputSegment = Segment{binary.LittleEndian.Uint64((*memory)[memLoc : memLoc+8]), binary.LittleEndian.Uint64((*memory)[memLoc+8 : memLoc+16])}
	memLoc += 16
	c.OutputSegment = Segment{binary.LittleEndian.Uint64((*memory)[memLoc : memLoc+8]), binary.LittleEndian.Uint64((*memory)[memLoc+8 : memLoc+16])}
	memLoc += 16
	c.SharedSegment = Segment{binary.LittleEndian.Uint64((*memory)[memLoc : memLoc+8]), binary.LittleEndian.Uint64((*memory)[memLoc+8 : memLoc+16])}
	memLoc += 16
	c.codePointer = binary.LittleEndian.Uint64((*memory)[memLoc : memLoc+8])
	memLoc += 8

}

func (e *Emulator) Step() {
	for _, c := range e.cores {
		c.Step()
	}
}

func (c *Core) Step() {
	if c.codePointer < c.startBlock || c.codePointer > c.startBlock+127 {
		// load new block
		newStartBlock := (c.codePointer / 256) * 256
		c.startBlock = newStartBlock
		c.codeCache = (*c.memory)[c.startBlock : c.startBlock+256]
	}

	opcode := c.codeCache[c.codePointer%256]
	c.execute(opcode)
}

func (c *Core) execute(opcode byte) {
	c.codePointer++
	mode := opcode >> 2
	if mode > 0 {
		if c.codePointer%256 == 0 {
			// bad opcode location
		} else {
			if mode == 1 {
				secondary := c.codeCache[c.codePointer%256 : c.codePointer%256+1]
				c.codePointer += 1
				c.DoOperation[opcode](secondary)
			} else if mode == 2 {
				secondary := c.codeCache[c.codePointer%256 : c.codePointer%256+2]
				c.codePointer += 2
				c.DoOperation[opcode](secondary)
			} else if mode == 3 {
				secondary := c.codeCache[c.codePointer%256 : c.codePointer%256+4]
				c.codePointer += 4
				c.DoOperation[opcode](secondary)
			}
		}
	} else {
		c.DoOperation[opcode](nil)
	}
}

func (c *Core) opADC(operand []byte) {
	accReg := int(c.status[0]) >> 4
	oprReg := int(c.status[0]) & 15

	carry := uint(0)
	sum := make([]byte, 32)
	for i := 31; i >= 0; i-- {
		tmp := uint(c.generalPurpose[accReg][i]) + uint(c.generalPurpose[oprReg][i]) + carry
		sum[i] = byte(tmp & 0xff)
		carry = tmp >> 8
	}

	c.generalPurpose[accReg] = sum

}
