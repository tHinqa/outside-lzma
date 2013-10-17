// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package lzma provides API definitions for accessing the liblzma dll.
package lzma

import (
	"github.com/tHinqa/outside"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type (
	Bool      bool
	Enum      int
	Vli       uint64
	Void      struct{}
	Internal  struct{}
	Index     struct{}
	IndexHash struct{}

	Stream struct {
		nextIn    *uint8
		availIn   uint
		totalIn   uint64
		nextOut   *uint8
		availOut  uint
		totalOut  uint64
		allocator *Allocator
		internal  *Internal
		_         *Void
		_         *Void
		_         *Void
		_         *Void
		_         uint64
		_         uint64
		_         uint
		_         uint
		_         Enum
		_         Enum
	}

	Allocator struct {
		Alloc  func(opaque *Void, nmemb uint, size uint) *Void
		Free   func(opaque *Void, ptr *Void)
		opaque *Void
	}

	Block struct {
		version          uint32
		headerSize       uint32
		check            Check
		compressedSize   Vli
		uncompressedSize Vli
		filters          *Filter
		rawCheck         [64]uint8
		_                *Void
		_                *Void
		_                *Void
		_                uint32
		_                uint32
		_                Vli
		_                Vli
		_                Vli
		_                Vli
		_                Vli
		_                Vli
		_                Enum
		_                Enum
		_                Enum
		_                Enum
		_                Bool
		_                Bool
		_                Bool
		_                Bool
		_                Bool
		_                Bool
		_                Bool
		_                Bool
	}

	Filter struct {
		id      Vli
		options *Void
	}

	OptionsLzma struct {
		dictSize       uint32
		presetDict     *uint8
		presetDictSize uint32
		lc             uint32
		lp             uint32
		pb             uint32
		mode           Mode
		niceLen        uint32
		mf             MatchFinder
		depth          uint32
		_              uint32
		_              uint32
		_              uint32
		_              uint32
		_              uint32
		_              uint32
		_              uint32
		_              uint32
		_              Enum
		_              Enum
		_              Enum
		_              Enum
		_              *Void
		_              *Void
	}

	IndexIter struct {
		Stream struct {
			flags              *StreamFlags
			_                  *Void
			_                  *Void
			_                  *Void
			number             Vli
			blockCount         Vli
			compressedOffset   Vli
			uncompressedOffset Vli
			compressedSize     Vli
			uncompressedSize   Vli
			padding            Vli
			_                  Vli
			_                  Vli
			_                  Vli
			_                  Vli
		}
		Block struct {
			numberInFile             Vli
			compressedFileOffset     Vli
			uncompressedFileOffset   Vli
			numberInStream           Vli
			compressedStreamOffset   Vli
			uncompressedStreamOffset Vli
			uncompressedSize         Vli
			unpaddedSize             Vli
			totalSize                Vli
			_                        Vli
			_                        Vli
			_                        Vli
			_                        Vli
			_                        *Void
			_                        *Void
			_                        *Void
			_                        *Void
		}
		internal [6]struct {
			//TODO(t): Union
			// p  *Void
			// s  uint
			v Vli
		}
	}

	StreamFlags struct {
		version      uint32
		backwardSize Vli
		check        Check
		_            Enum
		_            Enum
		_            Enum
		_            Enum
		_            Bool
		_            Bool
		_            Bool
		_            Bool
		_            Bool
		_            Bool
		_            Bool
		_            Bool
		_            uint32
		_            uint32
	}
)

var (
	Code func(strm *Stream, action Action) Ret

	End func(strm *Stream)

	Memusage func(strm *Stream) uint64

	MemlimitGet func(strm *Stream) uint64

	MemlimitSet func(strm *Stream, memlimit uint64) Ret

	VliEncode func(vli Vli, vliPos *uint,
		out *uint8, outPos *uint, outSize uint) Ret

	VliDecode func(vli *Vli, vliPos *uint,
		in *uint8, inPos *uint, inSize uint) Ret

	VliSize func(vli Vli) uint32

	CheckIsSupported func(check Check) Bool

	CheckSize func(check Check) uint32

	Crc32 func(buf *uint8, size uint, crc uint32) uint32

	Crc64 func(buf *uint8, size uint, crc uint64) uint64

	GetCheck func(strm *Stream) Check

	FilterEncoderIsSupported func(id Vli) Bool

	FilterDecoderIsSupported func(id Vli) Bool

	FiltersCopy func(src, dest *Filter, allocator *Allocator) Ret

	RawEncoderMemusage func(filters *Filter) uint64

	RawDecoderMemusage func(filters *Filter) uint64

	RawEncoder func(strm *Stream, filters *Filter) Ret

	RawDecoder func(strm *Stream, filters *Filter) Ret

	FiltersUpdate func(strm *Stream, filters *Filter) Ret

	RawBufferEncode func(filters *Filter, allocator *Allocator,
		in *uint8, inSize uint,
		out *uint8, outPos *uint, outSize uint) Ret

	RawBufferDecode func(filters *Filter, allocator *Allocator,
		in *uint8, inPos *uint, inSize uint,
		out *uint8, outPos *uint, outSize uint) Ret

	PropertiesSize func(size *uint32, filter *Filter) Ret

	PropertiesEncode func(filter *Filter, props *uint8) Ret

	PropertiesDecode func(filter *Filter, allocator *Allocator,
		props *uint8, propsSize uint) Ret

	FilterFlagsSize func(size *uint32, filter *Filter) Ret

	FilterFlagsEncode func(filter *Filter,
		out *uint8, outPos *uint, outSize uint) Ret

	FilterFlagsDecode func(filter *Filter, allocator *Allocator,
		in *uint8, inPos *uint, inSize uint) Ret

	MfIsSupported func(matchFinder MatchFinder) Bool

	ModeIsSupported func(mode Mode) Bool

	LzmaPreset func(options *OptionsLzma, preset uint32) Bool

	EasyEncoderMemusage func(preset uint32) uint64

	EasyDecoderMemusage func(preset uint32) uint64

	EasyEncoder func(strm *Stream, preset uint32, check Check) Ret

	EasyBufferEncode func(preset uint32, check Check,
		allocator *Allocator, in *uint8, inSize uint,
		out *uint8, outPos *uint, outSize uint) Ret

	StreamEncoder func(
		strm *Stream, filters *Filter, check Check) Ret

	AloneEncoder func(strm *Stream, options *OptionsLzma) Ret

	StreamBufferBound func(uncompressedSize uint) uint

	StreamBufferEncode func(filters *Filter, check Check,
		allocator *Allocator, in *uint8, inSize uint,
		out *uint8, outPos *uint, outSize uint) Ret

	StreamDecoder func(
		strm *Stream, memlimit uint64, flags uint32) Ret

	AutoDecoder func(
		strm *Stream, memlimit uint64, flags uint32) Ret

	AloneDecoder func(strm *Stream, memlimit uint64) Ret

	StreamBufferDecode func(memlimit *uint64, flags uint32,
		allocator *Allocator, in *uint8, inPos *uint, inSize uint,
		out *uint8, outPos *uint, outSize uint) Ret

	StreamHeaderEncode func(options *StreamFlags, out *uint8) Ret

	StreamFooterEncode func(options *StreamFlags, out *uint8) Ret

	StreamHeaderDecode func(options *StreamFlags, in *uint8) Ret

	StreamFooterDecode func(options *StreamFlags, in *uint8) Ret

	StreamFlagsCompare func(a, b *StreamFlags) Ret

	BlockHeaderSize func(block *Block) Ret

	BlockHeaderEncode func(block *Block, out *uint8) Ret

	BlockHeaderDecode func(
		block *Block, allocator *Allocator, in *uint8) Ret

	BlockCompressedSize func(block *Block, unpaddedSize Vli) Ret

	BlockUnpaddedSize func(block *Block) Vli

	BlockTotalSize func(block *Block) Vli

	BlockEncoder func(strm *Stream, block *Block) Ret

	BlockDecoder func(strm *Stream, block *Block) Ret

	BlockBufferBound func(uncompressedSize uint) uint

	BlockBufferEncode func(block *Block, allocator *Allocator,
		in *uint8, inSize uint,
		out *uint8, outPos *uint, outSize uint) Ret

	BlockBufferDecode func(block *Block, allocator *Allocator,
		in *uint8, inPos *uint, inSize uint,
		out *uint8, outPos *uint, outSize uint) Ret

	IndexMemusage func(streams Vli, blocks Vli) uint64

	IndexMemused func(i *Index) uint64

	IndexInit func(allocator *Allocator) *Index

	IndexEnd func(i *Index, allocator *Allocator)

	IndexAppend func(i *Index, allocator *Allocator,
		unpaddedSize, uncompressedSize Vli) Ret

	IndexStreamFlags func(i *Index, streamFlags *StreamFlags) Ret

	IndexChecks func(i *Index) uint32

	IndexStreamPadding func(i *Index, streamPadding Vli) Ret

	IndexStreamCount func(i *Index) Vli

	IndexBlockCount func(i *Index) Vli

	IndexSize func(i *Index) Vli

	IndexStreamSize func(i *Index) Vli

	IndexTotalSize func(i *Index) Vli

	IndexFileSize func(i *Index) Vli

	IndexUncompressedSize func(i *Index) Vli

	IndexIterInit func(iter *IndexIter, i *Index)

	IndexIterRewind func(iter *IndexIter)

	IndexIterNext func(iter *IndexIter, mode IndexIterMode) Bool

	IndexIterLocate func(iter *IndexIter, target Vli) Bool

	IndexCat func(dest, src *Index, allocator *Allocator) Ret

	IndexDup func(i *Index, allocator *Allocator) *Index

	IndexEncoder func(strm *Stream, i *Index) Ret

	IndexDecoder func(
		strm *Stream, i **Index, memlimit uint64) Ret

	IndexBufferEncode func(
		i *Index, out *uint8, outPos *uint, outSize uint) Ret

	IndexBufferDecode func(i **Index, memlimit *uint64,
		allocator *Allocator, in *uint8, inPos *uint,
		inSize uint) Ret

	IndexHashInit func(
		indexHash *IndexHash, allocator *Allocator) *IndexHash

	IndexHashEnd func(indexHash *IndexHash, allocator *Allocator)

	IndexHashAppend func(indexHash *IndexHash,
		unpaddedSize Vli, uncompressedSize Vli) Ret

	IndexHashDecode func(indexHash *IndexHash,
		in *uint8, inPos *uint, inSize uint) Ret

	IndexHashSize func(indexHash *IndexHash) Vli

	Physmem func() uint64

	VersionNumber func() uint32

	VersionString func() string
)

type Ret Enum

const (
	OK Ret = iota
	STREAM_END
	NO_CHECK
	UNSUPPORTED_CHECK
	GET_CHECK
	MEM_ERROR
	MEMLIMIT_ERROR
	FORMAT_ERROR
	OPTIONS_ERROR
	DATA_ERROR
	BUF_ERROR
	PROG_ERROR
)

type Check Enum

const (
	CHECK_NONE Check = iota
	CHECK_CRC32
	CHECK_CRC64  Check = 4
	CHECK_SHA256 Check = 10
)

type Action Enum

const (
	RUN Action = iota
	SYNC_FLUSH
	FULL_FLUSH
	FINISH
)

type IndexIterMode Enum

const (
	INDEX_ITER_ANY IndexIterMode = iota
	INDEX_ITER_STREAM
	INDEX_ITER_BLOCK
	INDEX_ITER_NONEMPTY_BLOCK
)

type MatchFinder Enum

const (
	MF_HC3 MatchFinder = iota + 0x03
	MF_HC4
)
const (
	MF_BT2 MatchFinder = iota + 0x12
	MF_BT3
	MF_BT4
)

type Mode Enum

const (
	MODE_FAST Mode = iota + 1
	MODE_NORMAL
)

var dll = "liblzma.dll"

var apiList = outside.Apis{
	{"lzma_alone_decoder", &AloneDecoder},
	{"lzma_alone_encoder", &AloneEncoder},
	{"lzma_auto_decoder", &AutoDecoder},
	{"lzma_block_buffer_bound", &BlockBufferBound},
	{"lzma_block_buffer_decode", &BlockBufferDecode},
	{"lzma_block_buffer_encode", &BlockBufferEncode},
	{"lzma_block_compressed_size", &BlockCompressedSize},
	{"lzma_block_decoder", &BlockDecoder},
	{"lzma_block_encoder", &BlockEncoder},
	{"lzma_block_header_decode", &BlockHeaderDecode},
	{"lzma_block_header_encode", &BlockHeaderEncode},
	{"lzma_block_header_size", &BlockHeaderSize},
	{"lzma_block_total_size", &BlockTotalSize},
	{"lzma_block_unpadded_size", &BlockUnpaddedSize},
	{"lzma_check_is_supported", &CheckIsSupported},
	{"lzma_check_size", &CheckSize},
	{"lzma_code", &Code},
	{"lzma_crc32", &Crc32},
	{"lzma_crc64", &Crc64},
	{"lzma_easy_buffer_encode", &EasyBufferEncode},
	{"lzma_easy_decoder_memusage", &EasyDecoderMemusage},
	{"lzma_easy_encoder", &EasyEncoder},
	{"lzma_easy_encoder_memusage", &EasyEncoderMemusage},
	{"lzma_end", &End},
	{"lzma_filter_decoder_is_supported", &FilterDecoderIsSupported},
	{"lzma_filter_encoder_is_supported", &FilterEncoderIsSupported},
	{"lzma_filter_flags_decode", &FilterFlagsDecode},
	{"lzma_filter_flags_encode", &FilterFlagsEncode},
	{"lzma_filter_flags_size", &FilterFlagsSize},
	{"lzma_filters_copy", &FiltersCopy},
	{"lzma_filters_update", &FiltersUpdate},
	{"lzma_get_check", &GetCheck},
	{"lzma_index_append", &IndexAppend},
	{"lzma_index_block_count", &IndexBlockCount},
	{"lzma_index_buffer_decode", &IndexBufferDecode},
	{"lzma_index_buffer_encode", &IndexBufferEncode},
	{"lzma_index_cat", &IndexCat},
	{"lzma_index_checks", &IndexChecks},
	{"lzma_index_decoder", &IndexDecoder},
	{"lzma_index_dup", &IndexDup},
	{"lzma_index_encoder", &IndexEncoder},
	{"lzma_index_end", &IndexEnd},
	{"lzma_index_file_size", &IndexFileSize},
	{"lzma_index_hash_append", &IndexHashAppend},
	{"lzma_index_hash_decode", &IndexHashDecode},
	{"lzma_index_hash_end", &IndexHashEnd},
	{"lzma_index_hash_init", &IndexHashInit},
	{"lzma_index_hash_size", &IndexHashSize},
	{"lzma_index_init", &IndexInit},
	{"lzma_index_iter_init", &IndexIterInit},
	{"lzma_index_iter_locate", &IndexIterLocate},
	{"lzma_index_iter_next", &IndexIterNext},
	{"lzma_index_iter_rewind", &IndexIterRewind},
	{"lzma_index_memusage", &IndexMemusage},
	{"lzma_index_memused", &IndexMemused},
	{"lzma_index_size", &IndexSize},
	{"lzma_index_stream_count", &IndexStreamCount},
	{"lzma_index_stream_flags", &IndexStreamFlags},
	{"lzma_index_stream_padding", &IndexStreamPadding},
	{"lzma_index_stream_size", &IndexStreamSize},
	{"lzma_index_total_size", &IndexTotalSize},
	{"lzma_index_uncompressed_size", &IndexUncompressedSize},
	{"lzma_lzma_preset", &LzmaPreset},
	{"lzma_memlimit_get", &MemlimitGet},
	{"lzma_memlimit_set", &MemlimitSet},
	{"lzma_memusage", &Memusage},
	{"lzma_mf_is_supported", &MfIsSupported},
	{"lzma_mode_is_supported", &ModeIsSupported},
	{"lzma_physmem", &Physmem},
	{"lzma_properties_decode", &PropertiesDecode},
	{"lzma_properties_encode", &PropertiesEncode},
	{"lzma_properties_size", &PropertiesSize},
	{"lzma_raw_buffer_decode", &RawBufferDecode},
	{"lzma_raw_buffer_encode", &RawBufferEncode},
	{"lzma_raw_decoder", &RawDecoder},
	{"lzma_raw_decoder_memusage", &RawDecoderMemusage},
	{"lzma_raw_encoder", &RawEncoder},
	{"lzma_raw_encoder_memusage", &RawEncoderMemusage},
	{"lzma_stream_buffer_bound", &StreamBufferBound},
	{"lzma_stream_buffer_decode", &StreamBufferDecode},
	{"lzma_stream_buffer_encode", &StreamBufferEncode},
	{"lzma_stream_decoder", &StreamDecoder},
	{"lzma_stream_encoder", &StreamEncoder},
	{"lzma_stream_flags_compare", &StreamFlagsCompare},
	{"lzma_stream_footer_decode", &StreamFooterDecode},
	{"lzma_stream_footer_encode", &StreamFooterEncode},
	{"lzma_stream_header_decode", &StreamHeaderDecode},
	{"lzma_stream_header_encode", &StreamHeaderEncode},
	{"lzma_version_number", &VersionNumber},
	{"lzma_version_string", &VersionString},
	{"lzma_vli_decode", &VliDecode},
	{"lzma_vli_encode", &VliEncode},
	{"lzma_vli_size", &VliSize},
}
