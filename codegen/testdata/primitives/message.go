package primitives

type SubMessage struct {

	A1 int32
}

type Message struct {
	Sub SubMessage
//	A1 int `parquet:"a1,omitempty"`
	//B1 uint
//	D1 int32 `parquet:"d1,omitempty"`
	//E1 uint32
//	G1 int64 `parquet:"g1,omitempty"`
	//H1 uint64
	//I1 bool
	//J1 string
	//K1 []byte
	//K1 float32
	//L1 float64
//	A1 []string
//	A2 []*int32
	//A2 *int32
	//A3 *int32
	//B2 *uint
	//D2 *int32
	//E2 *uint32
	//G2 *int64
	//H2 *uint64
	//I2 *bool
	//J2 *string
	//K2 *float32
	//L2 *float64
}
