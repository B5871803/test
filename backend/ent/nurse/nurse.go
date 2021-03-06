// Code generated by entc, DO NOT EDIT.

package nurse

const (
	// Label holds the string label denoting the nurse type in the database.
	Label = "nurse"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNurseName holds the string denoting the nurse_name field in the database.
	FieldNurseName = "nurse_name"
	// FieldNurseEmail holds the string denoting the nurse_email field in the database.
	FieldNurseEmail = "nurse_email"
	// FieldNursePassword holds the string denoting the nurse_password field in the database.
	FieldNursePassword = "nurse_password"
	// FieldNurseTel holds the string denoting the nurse_tel field in the database.
	FieldNurseTel = "nurse_tel"

	// EdgeNurseop holds the string denoting the nurseop edge name in mutations.
	EdgeNurseop = "nurseop"

	// Table holds the table name of the nurse in the database.
	Table = "nurses"
	// NurseopTable is the table the holds the nurseop relation/edge.
	NurseopTable = "operativerecords"
	// NurseopInverseTable is the table name for the Operativerecord entity.
	// It exists in this package in order to avoid circular dependency with the "operativerecord" package.
	NurseopInverseTable = "operativerecords"
	// NurseopColumn is the table column denoting the nurseop relation/edge.
	NurseopColumn = "nurse_id"
)

// Columns holds all SQL columns for nurse fields.
var Columns = []string{
	FieldID,
	FieldNurseName,
	FieldNurseEmail,
	FieldNursePassword,
	FieldNurseTel,
}

var (
	// NurseNameValidator is a validator for the "nurse_name" field. It is called by the builders before save.
	NurseNameValidator func(string) error
	// NurseEmailValidator is a validator for the "nurse_email" field. It is called by the builders before save.
	NurseEmailValidator func(string) error
	// NursePasswordValidator is a validator for the "nurse_password" field. It is called by the builders before save.
	NursePasswordValidator func(string) error
	// NurseTelValidator is a validator for the "nurse_tel" field. It is called by the builders before save.
	NurseTelValidator func(string) error
)
