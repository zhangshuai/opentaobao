package aesolution

// GlobalAeopTpLoanInfoDto 
type GlobalAeopTpLoanInfoDto struct {
    // loan amount
    LoanAmount   *GlobalMoneyStr `json:"loan_amount,omitempty" xml:"loan_amount,omitempty"`
    // loan time
    LoanTime   string `json:"loan_time,omitempty" xml:"loan_time,omitempty"`
}
