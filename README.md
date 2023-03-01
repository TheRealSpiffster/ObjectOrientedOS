# ObjectOrientedOS
An fully distributed parallel object oriented computing system.

## Problem:
  The OSes of today have been built over the years with the goal of compatibility to prior generations in mind.  While this allows a continuity of software, it disallows true innovation and security in the computing environment.

  Outposts on other planets will offer additional challenges.  Computing resources will be scarce, and communication with earth will be at a premium.  It is important to harness all the computing power available at the remote location.  Current OS makers (Linux, Microsoft, Apple, Google) build the OS to focus on the standalone user.  Security is an afterthought (literally there are 3rd party applications to enhance security).  These legacy systems do not share computational resources with other computers very well.  One would need to write specialized software in order to share computing cores across a network.

  Current OSes were never designed for today’s hardware and communications abilities.  Nor were they designed to conserve resources.  As I write this, most cores are idled wasting resources.  The hardware and OS is unable to allow cores to be used by remote software over networks in a seamless fashion.

## Concept Change:
  
  A new OS and CPU design could remove the old paradigm of “programs” and replaces it with “objects”.  The important difference between programs and objects is that programs cannot readily be used by other programs.  They are too big and complicated, and have only one entry point.  Objects, however, are designed to be used by other objects.  Once they are installed, anyone can make use of that object.

## Solution:

  A distributed OS requires a distributed design, and secure computing hardware.  The computers on a remote outpost, need an OS that joins them all as a super-computer to share the computing resources with the other scientists and colonists.  And share them in such a way that it is seamless and secure.

  The computers need to retain their individuality allowing them to run with or without the network of other computers.  They need to be able to prioritize computing for the owner, and only allowing idle time to be shared.
  
  The CPUs must be designed for security in the computing environment.  The CPU must present a sandbox for each object to run in.  The only mechanism for one object to interfere with another is a messaging system which an OS level controls.
  
  The CPU and OS must be designed for multi-threaded objects running on multiple cores in multiple computers.  The execution environment must isolate objects from other objects.
  
  The CPUs instruction set needs to be simplified to maximize the number of cores.  With the speed of the CPU far exceeding the speed of main memory, instructions should be imported in blocks to be executed without needing memory access on every instruction.  Using a 256 bit data bus with a 64 bit address bus, and an instruction size of 8 bits, doing two memory reads could bring in a block of 64 instructions into the core all at once.  Or this could be 128 instructions with 4 memory reads.  Code would execute in blocks.  Experimentation would have to resolve the best block size.
  
