package opcodes
import "strconv"
const _Opcode_name = "NopUnreachableSelectI32ConstI32AddI32SubI32MulI32DivSI32DivUI32RemSI32RemUI32AndI32OrI32XorI32ShlI32ShrSI32ShrUI32RotlI32RotrI32ClzI32CtzI32PopCntI32EqZI32EqI32NeI32LtSI32LtUI32LeSI32LeUI32GtSI32GtUI32GeSI32GeUI64ConstI64AddI64SubI64MulI64DivSI64DivUI64RemSI64RemUI64RotlI64RotrI64ClzI64CtzI64PopCntI64EqZI64AndI64OrI64XorI64ShlI64ShrSI64ShrUI64EqI64NeI64LtSI64LtUI64LeSI64LeUI64GtSI64GtUI64GeSI64GeUF32AddF32SubF32MulF32DivF32SqrtF32MinF32MaxF32CeilF32FloorF32TruncF32NearestF32AbsF32NegF32CopySignF32EqF32NeF32LtF32LeF32GtF32GeF64AddF64SubF64MulF64DivF64SqrtF64MinF64MaxF64CeilF64FloorF64TruncF64NearestF64AbsF64NegF64CopySignF64EqF64NeF64LtF64LeF64GtF64GeI32WrapI64I32TruncUF32I32TruncUF64I32TruncSF32I32TruncSF64I64TruncUF32I64TruncUF64I64TruncSF32I64TruncSF64I64ExtendUI32I64ExtendSI32F32DemoteF64F64PromoteF32F32ConvertSI32F32ConvertSI64F32ConvertUI32F32ConvertUI64F64ConvertSI32F64ConvertSI64F64ConvertUI32F64ConvertUI64I32LoadI64LoadI32StoreI64StoreI32Load8SI32Load16SI64Load8SI64Load16SI64Load32SI32Load8UI32Load16UI64Load8UI64Load16UI64Load32UI32Store8I32Store16I64Store8I64Store16I64Store32JmpJmpIfJmpEitherJmpTableReturnValueReturnVoidGetLocalSetLocalGetGlobalSetGlobalCallCallIndirectInvokeImportCurrentMemoryGrowMemoryPhiAddGasFPDisabledErrorUnknown"
var _Opcode_index = [...]uint16{0, 3, 14, 20, 28, 34, 40, 46, 53, 60, 67, 74, 80, 85, 91, 97, 104, 111, 118, 125, 131, 137, 146, 152, 157, 162, 168, 174, 180, 186, 192, 198, 204, 210, 218, 224, 230, 236, 243, 250, 257, 264, 271, 278, 284, 290, 299, 305, 311, 316, 322, 328, 335, 342, 347, 352, 358, 364, 370, 376, 382, 388, 394, 400, 406, 412, 418, 424, 431, 437, 443, 450, 458, 466, 476, 482, 488, 499, 504, 509, 514, 519, 524, 529, 535, 541, 547, 553, 560, 566, 572, 579, 587, 595, 605, 611, 617, 628, 633, 638, 643, 648, 653, 658, 668, 680, 692, 704, 716, 728, 740, 752, 764, 777, 790, 802, 815, 829, 843, 857, 871, 885, 899, 913, 927, 934, 941, 949, 957, 966, 976, 985, 995, 1005, 1014, 1024, 1033, 1043, 1053, 1062, 1072, 1081, 1091, 1101, 1104, 1109, 1118, 1126, 1137, 1147, 1155, 1163, 1172, 1181, 1185, 1197, 1209, 1222, 1232, 1235, 1241, 1256, 1263}

func (i Opcode) String() string {
	if i >= Opcode(len(_Opcode_index)-1) { return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")" }
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}