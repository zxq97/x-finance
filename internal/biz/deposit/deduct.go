package deposit

import (
	"context"

	"github.com/zxq97/x-finance/internal/biz"
)

func (uc *DepositUseCase) Deduct(ctx context.Context, param *DeductParam) error {
	main, err := uc.repo.GetMainOrderByID(ctx, param.MainID)
	if err != nil {
		return err
	}

	if main.FreezeType != 1 {
		return biz.ErrDepositNotFreeze
	}

	deducts, err := uc.repo.GetDeductsByMainID(ctx, param.MainID, -1)
	if err != nil {
		return err
	}

	balance := main.DepositAmount
	for _, v := range deducts {
		if v.Status == deducting || v.Status == deductFailed {
			return biz.ErrDeductPreCheckFailed
		} else if v.DeductNo == param.DeductNo {
			return nil
		} else if v.Status == deducted {
			balance -= v.Amount
		}
	}

	if balance < param.Amount {
		return biz.ErrDeductAmountInvalid
	}

	return nil
}
