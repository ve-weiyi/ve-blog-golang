package mqlogic

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/model"
)

// findDelivery 按幂等键查找投递记录；不存在返回 nil, nil。
func findDelivery(ctx context.Context, m model.TNotifyRecordModel, bizID, channel string) (*model.TNotifyRecord, error) {
	if bizID == "" {
		return nil, nil
	}

	record, err := m.FindOne(ctx, "", "biz_id = ? AND channel = ?", bizID, channel)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return record, nil
}

// createDelivery 插入投递记录；并发下唯一键冲突时返回已存在的记录。
func createDelivery(ctx context.Context, m model.TNotifyRecordModel, record *model.TNotifyRecord) (*model.TNotifyRecord, error) {
	if _, err := m.Insert(ctx, record); err != nil {
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, err
		}

		existing, findErr := findDelivery(ctx, m, record.BizId, record.Channel)
		if findErr != nil {
			return nil, findErr
		}
		if existing == nil {
			return nil, err
		}
		return existing, nil
	}
	return record, nil
}
