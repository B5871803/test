// Code generated by entc, DO NOT EDIT.

package operative

import (
	"github.com/B5871803/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OperativeType applies equality check predicate on the "operativeType" field. It's identical to OperativeTypeEQ.
func OperativeType(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperativeType), v))
	})
}

// OperativeName applies equality check predicate on the "operativeName" field. It's identical to OperativeNameEQ.
func OperativeName(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperativeName), v))
	})
}

// OperativeTypeEQ applies the EQ predicate on the "operativeType" field.
func OperativeTypeEQ(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeNEQ applies the NEQ predicate on the "operativeType" field.
func OperativeTypeNEQ(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeIn applies the In predicate on the "operativeType" field.
func OperativeTypeIn(vs ...string) predicate.Operative {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOperativeType), v...))
	})
}

// OperativeTypeNotIn applies the NotIn predicate on the "operativeType" field.
func OperativeTypeNotIn(vs ...string) predicate.Operative {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOperativeType), v...))
	})
}

// OperativeTypeGT applies the GT predicate on the "operativeType" field.
func OperativeTypeGT(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeGTE applies the GTE predicate on the "operativeType" field.
func OperativeTypeGTE(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeLT applies the LT predicate on the "operativeType" field.
func OperativeTypeLT(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeLTE applies the LTE predicate on the "operativeType" field.
func OperativeTypeLTE(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeContains applies the Contains predicate on the "operativeType" field.
func OperativeTypeContains(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeHasPrefix applies the HasPrefix predicate on the "operativeType" field.
func OperativeTypeHasPrefix(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeHasSuffix applies the HasSuffix predicate on the "operativeType" field.
func OperativeTypeHasSuffix(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeEqualFold applies the EqualFold predicate on the "operativeType" field.
func OperativeTypeEqualFold(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOperativeType), v))
	})
}

// OperativeTypeContainsFold applies the ContainsFold predicate on the "operativeType" field.
func OperativeTypeContainsFold(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOperativeType), v))
	})
}

// OperativeNameEQ applies the EQ predicate on the "operativeName" field.
func OperativeNameEQ(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperativeName), v))
	})
}

// OperativeNameNEQ applies the NEQ predicate on the "operativeName" field.
func OperativeNameNEQ(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOperativeName), v))
	})
}

// OperativeNameIn applies the In predicate on the "operativeName" field.
func OperativeNameIn(vs ...string) predicate.Operative {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOperativeName), v...))
	})
}

// OperativeNameNotIn applies the NotIn predicate on the "operativeName" field.
func OperativeNameNotIn(vs ...string) predicate.Operative {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOperativeName), v...))
	})
}

// OperativeNameGT applies the GT predicate on the "operativeName" field.
func OperativeNameGT(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOperativeName), v))
	})
}

// OperativeNameGTE applies the GTE predicate on the "operativeName" field.
func OperativeNameGTE(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOperativeName), v))
	})
}

// OperativeNameLT applies the LT predicate on the "operativeName" field.
func OperativeNameLT(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOperativeName), v))
	})
}

// OperativeNameLTE applies the LTE predicate on the "operativeName" field.
func OperativeNameLTE(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOperativeName), v))
	})
}

// OperativeNameContains applies the Contains predicate on the "operativeName" field.
func OperativeNameContains(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOperativeName), v))
	})
}

// OperativeNameHasPrefix applies the HasPrefix predicate on the "operativeName" field.
func OperativeNameHasPrefix(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOperativeName), v))
	})
}

// OperativeNameHasSuffix applies the HasSuffix predicate on the "operativeName" field.
func OperativeNameHasSuffix(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOperativeName), v))
	})
}

// OperativeNameEqualFold applies the EqualFold predicate on the "operativeName" field.
func OperativeNameEqualFold(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOperativeName), v))
	})
}

// OperativeNameContainsFold applies the ContainsFold predicate on the "operativeName" field.
func OperativeNameContainsFold(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOperativeName), v))
	})
}

// HasFromoperative applies the HasEdge predicate on the "fromoperative" edge.
func HasFromoperative() predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FromoperativeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FromoperativeTable, FromoperativeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFromoperativeWith applies the HasEdge predicate on the "fromoperative" edge with a given conditions (other predicates).
func HasFromoperativeWith(preds ...predicate.Operativerecord) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FromoperativeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FromoperativeTable, FromoperativeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Operative) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Operative) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Operative) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		p(s.Not())
	})
}
