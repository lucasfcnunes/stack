/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class WorkflowInstanceHistoryStageInput {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("ConfirmHold")public ActivityConfirmHold confirmHold;
    public WorkflowInstanceHistoryStageInput withConfirmHold(ActivityConfirmHold confirmHold) {
        this.confirmHold = confirmHold;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("CreateTransaction")public ActivityCreateTransaction createTransaction;
    public WorkflowInstanceHistoryStageInput withCreateTransaction(ActivityCreateTransaction createTransaction) {
        this.createTransaction = createTransaction;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("CreditWallet")public ActivityCreditWallet creditWallet;
    public WorkflowInstanceHistoryStageInput withCreditWallet(ActivityCreditWallet creditWallet) {
        this.creditWallet = creditWallet;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("DebitWallet")public ActivityDebitWallet debitWallet;
    public WorkflowInstanceHistoryStageInput withDebitWallet(ActivityDebitWallet debitWallet) {
        this.debitWallet = debitWallet;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("GetAccount")public ActivityGetAccount getAccount;
    public WorkflowInstanceHistoryStageInput withGetAccount(ActivityGetAccount getAccount) {
        this.getAccount = getAccount;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("GetPayment")public ActivityGetPayment getPayment;
    public WorkflowInstanceHistoryStageInput withGetPayment(ActivityGetPayment getPayment) {
        this.getPayment = getPayment;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("GetWallet")public ActivityGetWallet getWallet;
    public WorkflowInstanceHistoryStageInput withGetWallet(ActivityGetWallet getWallet) {
        this.getWallet = getWallet;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("RevertTransaction")public ActivityRevertTransaction revertTransaction;
    public WorkflowInstanceHistoryStageInput withRevertTransaction(ActivityRevertTransaction revertTransaction) {
        this.revertTransaction = revertTransaction;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("StripeTransfer")public ActivityStripeTransfer stripeTransfer;
    public WorkflowInstanceHistoryStageInput withStripeTransfer(ActivityStripeTransfer stripeTransfer) {
        this.stripeTransfer = stripeTransfer;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("VoidHold")public ActivityVoidHold voidHold;
    public WorkflowInstanceHistoryStageInput withVoidHold(ActivityVoidHold voidHold) {
        this.voidHold = voidHold;
        return this;
    }
    
}
