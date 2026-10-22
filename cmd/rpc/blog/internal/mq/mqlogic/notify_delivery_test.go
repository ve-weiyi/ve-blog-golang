package mqlogic

import (
	"context"
	"errors"
	"testing"

	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/model"
)

type fakeDeliveryModel struct {
	model.TNotifyRecordModel
	existing  *model.TNotifyRecord
	findErr   error
	insertErr error
	findCalls int
}

func (f *fakeDeliveryModel) FindOne(_ context.Context, _ string, _ string, _ ...interface{}) (*model.TNotifyRecord, error) {
	f.findCalls++
	if f.findErr != nil {
		return nil, f.findErr
	}
	if f.existing == nil {
		return nil, gorm.ErrRecordNotFound
	}
	return f.existing, nil
}

func (f *fakeDeliveryModel) Insert(_ context.Context, _ *model.TNotifyRecord) (int64, error) {
	return 0, f.insertErr
}

func TestFindDelivery(t *testing.T) {
	t.Run("empty biz id skips query", func(t *testing.T) {
		m := &fakeDeliveryModel{findErr: errors.New("should not be called")}
		got, err := findDelivery(context.Background(), m, "", "email")
		if err != nil || got != nil {
			t.Fatalf("findDelivery() = %v, %v; want nil, nil", got, err)
		}
		if m.findCalls != 0 {
			t.Fatalf("findCalls = %d; want 0", m.findCalls)
		}
	})

	t.Run("not found", func(t *testing.T) {
		got, err := findDelivery(context.Background(), &fakeDeliveryModel{}, "biz-1", "email")
		if err != nil || got != nil {
			t.Fatalf("findDelivery() = %v, %v; want nil, nil", got, err)
		}
	})

	t.Run("found", func(t *testing.T) {
		want := &model.TNotifyRecord{Id: 7}
		got, err := findDelivery(context.Background(), &fakeDeliveryModel{existing: want}, "biz-1", "email")
		if err != nil || got != want {
			t.Fatalf("findDelivery() = %v, %v; want %v, nil", got, err, want)
		}
	})
}

func TestCreateDelivery(t *testing.T) {
	t.Run("insert succeeds", func(t *testing.T) {
		in := &model.TNotifyRecord{BizId: "biz-1", Channel: "email"}
		got, err := createDelivery(context.Background(), &fakeDeliveryModel{}, in)
		if err != nil || got != in {
			t.Fatalf("createDelivery() = %v, %v; want %v, nil", got, err, in)
		}
	})

	t.Run("duplicate returns existing", func(t *testing.T) {
		existing := &model.TNotifyRecord{Id: 9, BizId: "biz-1", Channel: "email"}
		m := &fakeDeliveryModel{existing: existing, insertErr: gorm.ErrDuplicatedKey}
		got, err := createDelivery(context.Background(), m, &model.TNotifyRecord{BizId: "biz-1", Channel: "email"})
		if err != nil || got != existing {
			t.Fatalf("createDelivery() = %v, %v; want %v, nil", got, err, existing)
		}
	})

	t.Run("duplicate but not found returns error", func(t *testing.T) {
		m := &fakeDeliveryModel{insertErr: gorm.ErrDuplicatedKey}
		_, err := createDelivery(context.Background(), m, &model.TNotifyRecord{BizId: "biz-1", Channel: "email"})
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			t.Fatalf("createDelivery() error = %v; want duplicated key", err)
		}
	})

	t.Run("other insert error", func(t *testing.T) {
		boom := errors.New("boom")
		m := &fakeDeliveryModel{insertErr: boom}
		_, err := createDelivery(context.Background(), m, &model.TNotifyRecord{BizId: "biz-1", Channel: "email"})
		if !errors.Is(err, boom) {
			t.Fatalf("createDelivery() error = %v; want boom", err)
		}
	})
}
