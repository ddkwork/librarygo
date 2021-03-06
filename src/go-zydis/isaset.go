// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

// ISASet is an enum of Instruction Set Architectures.
type ISASet int

//go:generate stringer -type=ISASet -linecomment

// ISASet enum values.
const (
	ISASetInvalid                 ISASet = iota // INVALID
	ISASetADOX_ADCX                             // ADOX_ADCX
	ISASetAES                                   // AES
	ISASetAMD                                   // AMD
	ISASetAMD3DNOW                              // AMD3DNOW
	ISASetAMX_BF16                              // AMX_BF16
	ISASetAMX_INT8                              // AMX_INT8
	ISASetAMX_TILE                              // AMX_TILE
	ISASetAVX                                   // AVX
	ISASetAVX2                                  // AVX2
	ISASetAVX2GATHER                            // AVX2GATHER
	ISASetAVX512BW_128                          // AVX512BW_128
	ISASetAVX512BW_128N                         // AVX512BW_128N
	ISASetAVX512BW_256                          // AVX512BW_256
	ISASetAVX512BW_512                          // AVX512BW_512
	ISASetAVX512BW_KOP                          // AVX512BW_KOP
	ISASetAVX512CD_128                          // AVX512CD_128
	ISASetAVX512CD_256                          // AVX512CD_256
	ISASetAVX512CD_512                          // AVX512CD_512
	ISASetAVX512DQ_128                          // AVX512DQ_128
	ISASetAVX512DQ_128N                         // AVX512DQ_128N
	ISASetAVX512DQ_256                          // AVX512DQ_256
	ISASetAVX512DQ_512                          // AVX512DQ_512
	ISASetAVX512DQ_KOP                          // AVX512DQ_KOP
	ISASetAVX512DQ_SCALAR                       // AVX512DQ_SCALAR
	ISASetAVX512ER_512                          // AVX512ER_512
	ISASetAVX512ER_SCALAR                       // AVX512ER_SCALAR
	ISASetAVX512F_128                           // AVX512F_128
	ISASetAVX512F_128N                          // AVX512F_128N
	ISASetAVX512F_256                           // AVX512F_256
	ISASetAVX512F_512                           // AVX512F_512
	ISASetAVX512F_KOP                           // AVX512F_KOP
	ISASetAVX512F_SCALAR                        // AVX512F_SCALAR
	ISASetAVX512PF_512                          // AVX512PF_512
	ISASetAVX512_4FMAPS_512                     // AVX512_4FMAPS_512
	ISASetAVX512_4FMAPS_SCALAR                  // AVX512_4FMAPS_SCALAR
	ISASetAVX512_4VNNIW_512                     // AVX512_4VNNIW_512
	ISASetAVX512_BF16_128                       // AVX512_BF16_128
	ISASetAVX512_BF16_256                       // AVX512_BF16_256
	ISASetAVX512_BF16_512                       // AVX512_BF16_512
	ISASetAVX512_BITALG_128                     // AVX512_BITALG_128
	ISASetAVX512_BITALG_256                     // AVX512_BITALG_256
	ISASetAVX512_BITALG_512                     // AVX512_BITALG_512
	ISASetAVX512_GFNI_128                       // AVX512_GFNI_128
	ISASetAVX512_GFNI_256                       // AVX512_GFNI_256
	ISASetAVX512_GFNI_512                       // AVX512_GFNI_512
	ISASetAVX512_IFMA_128                       // AVX512_IFMA_128
	ISASetAVX512_IFMA_256                       // AVX512_IFMA_256
	ISASetAVX512_IFMA_512                       // AVX512_IFMA_512
	ISASetAVX512_VAES_128                       // AVX512_VAES_128
	ISASetAVX512_VAES_256                       // AVX512_VAES_256
	ISASetAVX512_VAES_512                       // AVX512_VAES_512
	ISASetAVX512_VBMI2_128                      // AVX512_VBMI2_128
	ISASetAVX512_VBMI2_256                      // AVX512_VBMI2_256
	ISASetAVX512_VBMI2_512                      // AVX512_VBMI2_512
	ISASetAVX512_VBMI_128                       // AVX512_VBMI_128
	ISASetAVX512_VBMI_256                       // AVX512_VBMI_256
	ISASetAVX512_VBMI_512                       // AVX512_VBMI_512
	ISASetAVX512_VNNI_128                       // AVX512_VNNI_128
	ISASetAVX512_VNNI_256                       // AVX512_VNNI_256
	ISASetAVX512_VNNI_512                       // AVX512_VNNI_512
	ISASetAVX512_VP2INTERSECT_128               // AVX512_VP2INTERSECT_128
	ISASetAVX512_VP2INTERSECT_256               // AVX512_VP2INTERSECT_256
	ISASetAVX512_VP2INTERSECT_512               // AVX512_VP2INTERSECT_512
	ISASetAVX512_VPCLMULQDQ_128                 // AVX512_VPCLMULQDQ_128
	ISASetAVX512_VPCLMULQDQ_256                 // AVX512_VPCLMULQDQ_256
	ISASetAVX512_VPCLMULQDQ_512                 // AVX512_VPCLMULQDQ_512
	ISASetAVX512_VPOPCNTDQ_128                  // AVX512_VPOPCNTDQ_128
	ISASetAVX512_VPOPCNTDQ_256                  // AVX512_VPOPCNTDQ_256
	ISASetAVX512_VPOPCNTDQ_512                  // AVX512_VPOPCNTDQ_512
	ISASetAVXAES                                // AVXAES
	ISASetAVX_GFNI                              // AVX_GFNI
	ISASetBMI1                                  // BMI1
	ISASetBMI2                                  // BMI2
	ISASetCET                                   // CET
	ISASetCLDEMOTE                              // CLDEMOTE
	ISASetCLFLUSHOPT                            // CLFLUSHOPT
	ISASetCLFSH                                 // CLFSH
	ISASetCLWB                                  // CLWB
	ISASetCLZERO                                // CLZERO
	ISASetCMOV                                  // CMOV
	ISASetCMPXCHG16B                            // CMPXCHG16B
	ISASetENQCMD                                // ENQCMD
	ISASetF16C                                  // F16C
	ISASetFAT_NOP                               // FAT_NOP
	ISASetFCMOV                                 // FCMOV
	ISASetFMA                                   // FMA
	ISASetFMA4                                  // FMA4
	ISASetFXSAVE                                // FXSAVE
	ISASetFXSAVE64                              // FXSAVE64
	ISASetGFNI                                  // GFNI
	ISASetI186                                  // I186
	ISASetI286PROTECTED                         // I286PROTECTED
	ISASetI286REAL                              // I286REAL
	ISASetI386                                  // I386
	ISASetI486                                  // I486
	ISASetI486REAL                              // I486REAL
	ISASetI86                                   // I86
	ISASetINVPCID                               // INVPCID
	ISASetKNCE                                  // KNCE
	ISASetKNCJKBR                               // KNCJKBR
	ISASetKNCSTREAM                             // KNCSTREAM
	ISASetKNCV                                  // KNCV
	ISASetKNC_MISC                              // KNC_MISC
	ISASetKNC_PF_HINT                           // KNC_PF_HINT
	ISASetLAHF                                  // LAHF
	ISASetLONGMODE                              // LONGMODE
	ISASetLZCNT                                 // LZCNT
	ISASetMCOMMIT                               // MCOMMIT
	ISASetMONITOR                               // MONITOR
	ISASetMONITORX                              // MONITORX
	ISASetMOVBE                                 // MOVBE
	ISASetMOVDIR                                // MOVDIR
	ISASetMPX                                   // MPX
	ISASetPADLOCK_ACE                           // PADLOCK_ACE
	ISASetPADLOCK_PHE                           // PADLOCK_PHE
	ISASetPADLOCK_PMM                           // PADLOCK_PMM
	ISASetPADLOCK_RNG                           // PADLOCK_RNG
	ISASetPAUSE                                 // PAUSE
	ISASetPCLMULQDQ                             // PCLMULQDQ
	ISASetPCONFIG                               // PCONFIG
	ISASetPENTIUMMMX                            // PENTIUMMMX
	ISASetPENTIUMREAL                           // PENTIUMREAL
	ISASetPKU                                   // PKU
	ISASetPOPCNT                                // POPCNT
	ISASetPPRO                                  // PPRO
	ISASetPREFETCHWT1                           // PREFETCHWT1
	ISASetPREFETCH_NOP                          // PREFETCH_NOP
	ISASetPT                                    // PT
	ISASetRDPID                                 // RDPID
	ISASetRDPMC                                 // RDPMC
	ISASetRDPRU                                 // RDPRU
	ISASetRDRAND                                // RDRAND
	ISASetRDSEED                                // RDSEED
	ISASetRDTSCP                                // RDTSCP
	ISASetRDWRFSGS                              // RDWRFSGS
	ISASetRTM                                   // RTM
	ISASetSERIALIZE                             // SERIALIZE
	ISASetSGX                                   // SGX
	ISASetSGX_ENCLV                             // SGX_ENCLV
	ISASetSHA                                   // SHA
	ISASetSMAP                                  // SMAP
	ISASetSMX                                   // SMX
	ISASetSSE                                   // SSE
	ISASetSSE2                                  // SSE2
	ISASetSSE2MMX                               // SSE2MMX
	ISASetSSE3                                  // SSE3
	ISASetSSE3X87                               // SSE3X87
	ISASetSSE4                                  // SSE4
	ISASetSSE42                                 // SSE42
	ISASetSSE4A                                 // SSE4A
	ISASetSSEMXCSR                              // SSEMXCSR
	ISASetSSE_PREFETCH                          // SSE_PREFETCH
	ISASetSSSE3                                 // SSSE3
	ISASetSSSE3MMX                              // SSSE3MMX
	ISASetSVM                                   // SVM
	ISASetTBM                                   // TBM
	ISASetTSX_LDTRK                             // TSX_LDTRK
	ISASetVAES                                  // VAES
	ISASetVMFUNC                                // VMFUNC
	ISASetVPCLMULQDQ                            // VPCLMULQDQ
	ISASetVTX                                   // VTX
	ISASetWAITPKG                               // WAITPKG
	ISASetX87                                   // X87
	ISASetXOP                                   // XOP
	ISASetXSAVE                                 // XSAVE
	ISASetXSAVEC                                // XSAVEC
	ISASetXSAVEOPT                              // XSAVEOPT
	ISASetXSAVES                                // XSAVES
)
