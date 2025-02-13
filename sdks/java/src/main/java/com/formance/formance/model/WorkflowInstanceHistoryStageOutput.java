/*
 * Formance Stack API
 * Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 
 *
 * The version of the OpenAPI document: develop
 * Contact: support@formance.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.formance.formance.model;

import java.util.Objects;
import java.util.Arrays;
import com.formance.formance.model.AccountResponse;
import com.formance.formance.model.DebitWalletResponse;
import com.formance.formance.model.GetWalletResponse;
import com.formance.formance.model.PaymentResponse;
import com.formance.formance.model.TransactionResponse;
import com.formance.formance.model.TransactionsResponse;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/**
 * WorkflowInstanceHistoryStageOutput
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class WorkflowInstanceHistoryStageOutput {
  public static final String SERIALIZED_NAME_GET_ACCOUNT = "GetAccount";
  @SerializedName(SERIALIZED_NAME_GET_ACCOUNT)
  private AccountResponse getAccount;

  public static final String SERIALIZED_NAME_CREATE_TRANSACTION = "CreateTransaction";
  @SerializedName(SERIALIZED_NAME_CREATE_TRANSACTION)
  private TransactionsResponse createTransaction;

  public static final String SERIALIZED_NAME_REVERT_TRANSACTION = "RevertTransaction";
  @SerializedName(SERIALIZED_NAME_REVERT_TRANSACTION)
  private TransactionResponse revertTransaction;

  public static final String SERIALIZED_NAME_GET_PAYMENT = "GetPayment";
  @SerializedName(SERIALIZED_NAME_GET_PAYMENT)
  private PaymentResponse getPayment;

  public static final String SERIALIZED_NAME_DEBIT_WALLET = "DebitWallet";
  @SerializedName(SERIALIZED_NAME_DEBIT_WALLET)
  private DebitWalletResponse debitWallet;

  public static final String SERIALIZED_NAME_GET_WALLET = "GetWallet";
  @SerializedName(SERIALIZED_NAME_GET_WALLET)
  private GetWalletResponse getWallet;

  public WorkflowInstanceHistoryStageOutput() {
  }

  public WorkflowInstanceHistoryStageOutput getAccount(AccountResponse getAccount) {
    
    this.getAccount = getAccount;
    return this;
  }

   /**
   * Get getAccount
   * @return getAccount
  **/
  @javax.annotation.Nullable

  public AccountResponse getGetAccount() {
    return getAccount;
  }


  public void setGetAccount(AccountResponse getAccount) {
    this.getAccount = getAccount;
  }


  public WorkflowInstanceHistoryStageOutput createTransaction(TransactionsResponse createTransaction) {
    
    this.createTransaction = createTransaction;
    return this;
  }

   /**
   * Get createTransaction
   * @return createTransaction
  **/
  @javax.annotation.Nullable

  public TransactionsResponse getCreateTransaction() {
    return createTransaction;
  }


  public void setCreateTransaction(TransactionsResponse createTransaction) {
    this.createTransaction = createTransaction;
  }


  public WorkflowInstanceHistoryStageOutput revertTransaction(TransactionResponse revertTransaction) {
    
    this.revertTransaction = revertTransaction;
    return this;
  }

   /**
   * Get revertTransaction
   * @return revertTransaction
  **/
  @javax.annotation.Nullable

  public TransactionResponse getRevertTransaction() {
    return revertTransaction;
  }


  public void setRevertTransaction(TransactionResponse revertTransaction) {
    this.revertTransaction = revertTransaction;
  }


  public WorkflowInstanceHistoryStageOutput getPayment(PaymentResponse getPayment) {
    
    this.getPayment = getPayment;
    return this;
  }

   /**
   * Get getPayment
   * @return getPayment
  **/
  @javax.annotation.Nullable

  public PaymentResponse getGetPayment() {
    return getPayment;
  }


  public void setGetPayment(PaymentResponse getPayment) {
    this.getPayment = getPayment;
  }


  public WorkflowInstanceHistoryStageOutput debitWallet(DebitWalletResponse debitWallet) {
    
    this.debitWallet = debitWallet;
    return this;
  }

   /**
   * Get debitWallet
   * @return debitWallet
  **/
  @javax.annotation.Nullable

  public DebitWalletResponse getDebitWallet() {
    return debitWallet;
  }


  public void setDebitWallet(DebitWalletResponse debitWallet) {
    this.debitWallet = debitWallet;
  }


  public WorkflowInstanceHistoryStageOutput getWallet(GetWalletResponse getWallet) {
    
    this.getWallet = getWallet;
    return this;
  }

   /**
   * Get getWallet
   * @return getWallet
  **/
  @javax.annotation.Nullable

  public GetWalletResponse getGetWallet() {
    return getWallet;
  }


  public void setGetWallet(GetWalletResponse getWallet) {
    this.getWallet = getWallet;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    WorkflowInstanceHistoryStageOutput workflowInstanceHistoryStageOutput = (WorkflowInstanceHistoryStageOutput) o;
    return Objects.equals(this.getAccount, workflowInstanceHistoryStageOutput.getAccount) &&
        Objects.equals(this.createTransaction, workflowInstanceHistoryStageOutput.createTransaction) &&
        Objects.equals(this.revertTransaction, workflowInstanceHistoryStageOutput.revertTransaction) &&
        Objects.equals(this.getPayment, workflowInstanceHistoryStageOutput.getPayment) &&
        Objects.equals(this.debitWallet, workflowInstanceHistoryStageOutput.debitWallet) &&
        Objects.equals(this.getWallet, workflowInstanceHistoryStageOutput.getWallet);
  }

  @Override
  public int hashCode() {
    return Objects.hash(getAccount, createTransaction, revertTransaction, getPayment, debitWallet, getWallet);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WorkflowInstanceHistoryStageOutput {\n");
    sb.append("    getAccount: ").append(toIndentedString(getAccount)).append("\n");
    sb.append("    createTransaction: ").append(toIndentedString(createTransaction)).append("\n");
    sb.append("    revertTransaction: ").append(toIndentedString(revertTransaction)).append("\n");
    sb.append("    getPayment: ").append(toIndentedString(getPayment)).append("\n");
    sb.append("    debitWallet: ").append(toIndentedString(debitWallet)).append("\n");
    sb.append("    getWallet: ").append(toIndentedString(getWallet)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

