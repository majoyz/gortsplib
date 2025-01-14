package rtph264

// NALUType is the type of a NALU.
type NALUType uint8

// standard NALU types.
const (
	NALUTypeNonIDR                        NALUType = 1
	NALUTypeDataPartitionA                NALUType = 2
	NALUTypeDataPartitionB                NALUType = 3
	NALUTypeDataPartitionC                NALUType = 4
	NALUTypeIDR                           NALUType = 5
	NALUTypeSei                           NALUType = 6
	NALUTypeSPS                           NALUType = 7
	NALUTypePPS                           NALUType = 8
	NALUTypeAccessUnitDelimiter           NALUType = 9
	NALUTypeEndOfSequence                 NALUType = 10
	NALUTypeEndOfStream                   NALUType = 11
	NALUTypeFillerData                    NALUType = 12
	NALUTypeSPSExtension                  NALUType = 13
	NALUTypePrefix                        NALUType = 14
	NALUTypeSubsetSPS                     NALUType = 15
	NALUTypeReserved16                    NALUType = 16
	NALUTypeReserved17                    NALUType = 17
	NALUTypeReserved18                    NALUType = 18
	NALUTypeSliceLayerWithoutPartitioning NALUType = 19
	NALUTypeSliceExtension                NALUType = 20
	NALUTypeSliceExtensionDepth           NALUType = 21
	NALUTypeReserved22                    NALUType = 22
	NALUTypeReserved23                    NALUType = 23
	NALUTypeStapA                         NALUType = 24
	NALUTypeStapB                         NALUType = 25
	NALUTypeMtap16                        NALUType = 26
	NALUTypeMtap24                        NALUType = 27
	NALUTypeFuA                           NALUType = 28
	NALUTypeFuB                           NALUType = 29
)

// String implements fmt.Stringer.
func (nt NALUType) String() string {
	switch nt {
	case NALUTypeNonIDR:
		return "NonIDR"
	case NALUTypeDataPartitionA:
		return "DataPartitionA"
	case NALUTypeDataPartitionB:
		return "DataPartitionB"
	case NALUTypeDataPartitionC:
		return "DataPartitionC"
	case NALUTypeIDR:
		return "IDR"
	case NALUTypeSei:
		return "Sei"
	case NALUTypeSPS:
		return "SPS"
	case NALUTypePPS:
		return "PPS"
	case NALUTypeAccessUnitDelimiter:
		return "AccessUnitDelimiter"
	case NALUTypeEndOfSequence:
		return "EndOfSequence"
	case NALUTypeEndOfStream:
		return "EndOfStream"
	case NALUTypeFillerData:
		return "FillerData"
	case NALUTypeSPSExtension:
		return "SPSExtension"
	case NALUTypePrefix:
		return "Prefix"
	case NALUTypeSubsetSPS:
		return "SubsetSPS"
	case NALUTypeReserved16:
		return "Reserved16"
	case NALUTypeReserved17:
		return "Reserved17"
	case NALUTypeReserved18:
		return "Reserved18"
	case NALUTypeSliceLayerWithoutPartitioning:
		return "SliceLayerWithoutPartitioning"
	case NALUTypeSliceExtension:
		return "SliceExtension"
	case NALUTypeSliceExtensionDepth:
		return "SliceExtensionDepth"
	case NALUTypeReserved22:
		return "Reserved22"
	case NALUTypeReserved23:
		return "Reserved23"
	case NALUTypeStapA:
		return "StapA"
	case NALUTypeStapB:
		return "StapB"
	case NALUTypeMtap16:
		return "Mtap16"
	case NALUTypeMtap24:
		return "Mtap24"
	case NALUTypeFuA:
		return "FuA"
	case NALUTypeFuB:
		return "FuB"
	}
	return "unknown"
}
