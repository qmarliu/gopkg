package ethutils

import "math/big"

func GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int,
	swapNumerator *big.Int, swapDenominator *big.Int) (amountOut *big.Int) {
	amountInWithFee := big.NewInt(0)
	numerator := big.NewInt(0)
	denominator := big.NewInt(0)
	amountOut = big.NewInt(0)

	amountInWithFee.Mul(amountIn, swapNumerator)
	numerator.Mul(amountInWithFee, reserveOut)
	denominator.Mul(reserveIn, swapDenominator)
	denominator.Add(denominator, amountInWithFee)
	amountOut.Div(numerator, denominator)
	return
}

func GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int,
	swapNumerator *big.Int, swapDenominator *big.Int) (amountIn *big.Int) {
	numerator := big.NewInt(0)
	denominator := big.NewInt(0)
	amountIn = big.NewInt(0)

	numerator.Mul(reserveIn, amountOut)
	numerator.Mul(numerator, swapDenominator)

	denominator.Sub(reserveOut, amountOut)
	denominator.Mul(denominator, swapNumerator)
	amountIn.Div(numerator, denominator)
	amountIn.Add(amountIn, big.NewInt(1))
	return
}
