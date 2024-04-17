
// 查询{{.upperStartCamelObject}}记录
func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, id int64) (out *{{.upperStartCamelObject}}, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}
