module github.com/sub2pewds12/E1AP

go 1.25.5

require (
	github.com/ishidawataru/sctp v0.0.0-20230406120618-7ff4192f6ff2
	github.com/lvdund/asn1go v0.0.0-00010101000000-000000000000
	github.com/lvdund/ngap v1.4.13
)

require github.com/reogac/utils v1.0.0 // indirect

replace github.com/lvdund/asn1go => ./asn1go-tmp
