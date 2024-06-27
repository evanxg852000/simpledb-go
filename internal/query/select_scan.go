package query

// The scan class corresponding to the <i>select</i> relational
// algebra operator.
// All methods except next delegate their work to the underlying scan
type SelectScan struct {
	scan      Scan
	predicate *Predicate
}

// Create a select scan having the specified underlying
// scan and predicate.
func NewSelectScan(scan Scan, predicate *Predicate) *SelectScan {
	return &SelectScan{scan, predicate}
}

// Scan methods

func (ss *SelectScan) BeforeFirst() {
	ss.scan.BeforeFirst()
}

func (ss *SelectScan) Next() bool {
	for ss.scan.Next() {
		if ss.predicate.IsSatisfied(ss.scan) {
			return true
		}
	}
	return false
}

func (ss *SelectScan) GetInt(fieldName string) int64 {
	return ss.scan.GetInt(fieldName)
}

func (ss *SelectScan) GetString(fieldName string) string {
	return ss.scan.GetString(fieldName)
}

func (ss *SelectScan) GetValue(fieldName string) Constant {
	return ss.scan.GetValue(fieldName)
}

func (ss *SelectScan) HasField(fieldName string) bool {
	return ss.scan.HasField(fieldName)
}

func (ss *SelectScan) Close() {
	ss.scan.Close()
}

// UpdateScan methods

func (ss *SelectScan) SetInt(fieldName string, value int64) {
	us := any(ss.scan).(UpdateScan)
	us.SetInt(fieldName, value)
}

func (ss *SelectScan) SetString(fieldName string, value string) {
	us := any(ss.scan).(UpdateScan)
	us.SetString(fieldName, value)
}

func (ss *SelectScan) SetValue(fieldName string, value Constant) {
	us := any(ss.scan).(UpdateScan)
	us.SetValue(fieldName, value)
}

func (ss *SelectScan) Delete() {
	us := any(ss.scan).(UpdateScan)
	us.Delete()
}

func (ss *SelectScan) Insert() {
	us := any(ss.scan).(UpdateScan)
	us.Insert()
}

func (ss *SelectScan) GetRID() RID {
	us := any(ss.scan).(UpdateScan)
	return us.GetRID()
}

func (ss *SelectScan) MoveToRID(rId RID) {
	us := any(ss.scan).(UpdateScan)
	us.MoveToRID(rId)
}
