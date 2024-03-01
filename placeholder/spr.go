package placeholder

import (
	wof_spr "github.com/whosonfirst/go-whosonfirst-spr/v2"
	"github.com/sfomuseum/go-edtf"
	"github.com/whosonfirst/go-whosonfirst-flags"	
)

type PlaceholderRecordSPR struct {
	wof_spr.SPR
	record *results.PlaceholderRecord
}

type PlaceholderStandardPlacesResults struct {
	results []*PlaceholderRecordSPR
}

func (r *PlaceholderStandardPlacesResults) Results() []StandardPlacesResult {
	return r.results
}

func NewPlaceholderRecordSPR(r *results.PlaceholderRecord) (wof_spr.SPR, error) {
	
	s := &PlaceholderRecordSPR{
		record: r,
	}
	
	return s, nil
}

func (s *PlaceholderRecordSPR) Id() string {

}

func (s *PlaceholderRecordSPR) ParentId() string {

}

func (s *PlaceholderRecordSPR) Name() string {

}

func (s *PlaceholderRecordSPR) Placetype() string {

}

func (s *PlaceholderRecordSPR) Country() string {

}

func (s *PlaceholderRecordSPR) Repo() string {

}


func (s *PlaceholderRecordSPR) Path() string {

}

func (s *PlaceholderRecordSPR) URI() string {

}

func (s *PlaceholderRecordSPR) Inception() *edtf.EDTFDate {

}

func (s *PlaceholderRecordSPR) Cessation() *edtf.EDTFDate {

}

func (s *PlaceholderRecordSPR) Latitude() float64 {

}

func (s *PlaceholderRecordSPR) Longitude() float64 {

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

}

func (s *PlaceholderRecordSPR) Supersedes() []int64 {

}

func (s *PlaceholderRecordSPR) BelongsTo() []int64 {

}

func (s *PlaceholderRecordSPR) LastModified() int64 {

}


