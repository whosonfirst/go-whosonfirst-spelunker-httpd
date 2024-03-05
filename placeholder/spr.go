package placeholder

import (
	wof_spr "github.com/whosonfirst/go-whosonfirst-spr/v2"
	"github.com/sfomuseum/go-edtf"
	"github.com/whosonfirst/go-whosonfirst-flags"	
)

type PlaceholderRecordSPR struct {
	wof_spr.SPR
	record *results.PlaceholderRecord
	bbox [4]float64
}

type PlaceholderStandardPlacesResults struct {
	results []*PlaceholderRecordSPR
}

func (r *PlaceholderStandardPlacesResults) Results() []StandardPlacesResult {
	return r.results
}

func NewPlaceholderRecordSPR(r *results.PlaceholderRecord) (wof_spr.SPR, error) {

	str_bbox := strings.Split(r.Geometry.BoundingBox, ",")
	
	min_lon, err := strconv.ParseFloat64(str_bbox[0], 10)
	min_lat, err := strconv.ParseFloat64(str_bbox[1], 10)
	max_lon, err := strconv.ParseFloat64(str_bbox[2], 10)
	max_lat, err := strconv.ParseFloat64(str_bbox[3], 10)	
	
	
	s := &PlaceholderRecordSPR{
		record: r,
	}
	
	return s, nil
}

func (s *PlaceholderRecordSPR) Id() string {
	return strconv.FormatInt(s.record.Id, 10)
}

func (s *PlaceholderRecordSPR) ParentId() string {
	return "-1"
}

func (s *PlaceholderRecordSPR) Name() string {
	return s.record.Name
}

func (s *PlaceholderRecordSPR) Placetype() string {
	return s.record.Placetype
}

func (s *PlaceholderRecordSPR) Country() string {
	return ""
}

func (s *PlaceholderRecordSPR) Repo() string {
	return ""
}


func (s *PlaceholderRecordSPR) Path() string {
	return ""
}

func (s *PlaceholderRecordSPR) URI() string {
	return ""
}

func (s *PlaceholderRecordSPR) Inception() *edtf.EDTFDate {
	return s.unknownEDTF()	
}

func (s *PlaceholderRecordSPR) Cessation() *edtf.EDTFDate {
	return s.unknownEDTF()
}

func (s *PlaceholderRecordSPR) unknownEDTF() *edtf.EDTFDate {
	
	sp := common.UnknownDateSpan()
	
	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf.UNKNOWN,
		Level:   -1,
		Feature: "Unknown",
	}
	
	return d
}

func (s *PlaceholderRecordSPR) Latitude() float64 {
	return s.record.Geometry.Latitude
}

func (s *PlaceholderRecordSPR) Longitude() float64 {
	return s.record.Geometry.Longitude
}

func (s *PlaceholderRecordSPR) MinLatitude() float64 {

}

func (s *PlaceholderRecordSPR) MinLongitude() float64 {

}

func (s *PlaceholderRecordSPR) MaxLatitude() float64 {

}

func (s *PlaceholderRecordSPR) MaxLongitude() float64 {

}

func (s *PlaceholderRecordSPR) IsCurrent() flags.ExistentialFlag {

}

func (s *PlaceholderRecordSPR) IsCeased() flags.ExistentialFlag {

}

func (s *PlaceholderRecordSPR) IsDeprecated() flags.ExistentialFlag {

}

func (s *PlaceholderRecordSPR) IsSuperseded() flags.ExistentialFlag {

}

func (s *PlaceholderRecordSPR) IsSuperseding() flags.ExistentialFlag {

}

func (s *PlaceholderRecordSPR) SupersededBy() []int64 {
	return make([]int64, 0)
}

func (s *PlaceholderRecordSPR) Supersedes() []int64 {
	return make([]int64, 0)
}

func (s *PlaceholderRecordSPR) BelongsTo() []int64 {
	return make([]int64, 0)
}

func (s *PlaceholderRecordSPR) LastModified() int64 {
	return -1
}


