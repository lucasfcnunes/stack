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
import com.formance.formance.model.Transaction;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * TransactionsCursorResponseCursor
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class TransactionsCursorResponseCursor {
  public static final String SERIALIZED_NAME_PAGE_SIZE = "pageSize";
  @SerializedName(SERIALIZED_NAME_PAGE_SIZE)
  private Long pageSize;

  public static final String SERIALIZED_NAME_HAS_MORE = "hasMore";
  @SerializedName(SERIALIZED_NAME_HAS_MORE)
  private Boolean hasMore;

  public static final String SERIALIZED_NAME_PREVIOUS = "previous";
  @SerializedName(SERIALIZED_NAME_PREVIOUS)
  private String previous;

  public static final String SERIALIZED_NAME_NEXT = "next";
  @SerializedName(SERIALIZED_NAME_NEXT)
  private String next;

  public static final String SERIALIZED_NAME_DATA = "data";
  @SerializedName(SERIALIZED_NAME_DATA)
  private List<Transaction> data = new ArrayList<>();

  public TransactionsCursorResponseCursor() {
  }

  public TransactionsCursorResponseCursor pageSize(Long pageSize) {
    
    this.pageSize = pageSize;
    return this;
  }

   /**
   * Get pageSize
   * minimum: 1
   * maximum: 1000
   * @return pageSize
  **/
  @javax.annotation.Nonnull

  public Long getPageSize() {
    return pageSize;
  }


  public void setPageSize(Long pageSize) {
    this.pageSize = pageSize;
  }


  public TransactionsCursorResponseCursor hasMore(Boolean hasMore) {
    
    this.hasMore = hasMore;
    return this;
  }

   /**
   * Get hasMore
   * @return hasMore
  **/
  @javax.annotation.Nonnull

  public Boolean getHasMore() {
    return hasMore;
  }


  public void setHasMore(Boolean hasMore) {
    this.hasMore = hasMore;
  }


  public TransactionsCursorResponseCursor previous(String previous) {
    
    this.previous = previous;
    return this;
  }

   /**
   * Get previous
   * @return previous
  **/
  @javax.annotation.Nullable

  public String getPrevious() {
    return previous;
  }


  public void setPrevious(String previous) {
    this.previous = previous;
  }


  public TransactionsCursorResponseCursor next(String next) {
    
    this.next = next;
    return this;
  }

   /**
   * Get next
   * @return next
  **/
  @javax.annotation.Nullable

  public String getNext() {
    return next;
  }


  public void setNext(String next) {
    this.next = next;
  }


  public TransactionsCursorResponseCursor data(List<Transaction> data) {
    
    this.data = data;
    return this;
  }

  public TransactionsCursorResponseCursor addDataItem(Transaction dataItem) {
    this.data.add(dataItem);
    return this;
  }

   /**
   * Get data
   * @return data
  **/
  @javax.annotation.Nonnull

  public List<Transaction> getData() {
    return data;
  }


  public void setData(List<Transaction> data) {
    this.data = data;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TransactionsCursorResponseCursor transactionsCursorResponseCursor = (TransactionsCursorResponseCursor) o;
    return Objects.equals(this.pageSize, transactionsCursorResponseCursor.pageSize) &&
        Objects.equals(this.hasMore, transactionsCursorResponseCursor.hasMore) &&
        Objects.equals(this.previous, transactionsCursorResponseCursor.previous) &&
        Objects.equals(this.next, transactionsCursorResponseCursor.next) &&
        Objects.equals(this.data, transactionsCursorResponseCursor.data);
  }

  @Override
  public int hashCode() {
    return Objects.hash(pageSize, hasMore, previous, next, data);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TransactionsCursorResponseCursor {\n");
    sb.append("    pageSize: ").append(toIndentedString(pageSize)).append("\n");
    sb.append("    hasMore: ").append(toIndentedString(hasMore)).append("\n");
    sb.append("    previous: ").append(toIndentedString(previous)).append("\n");
    sb.append("    next: ").append(toIndentedString(next)).append("\n");
    sb.append("    data: ").append(toIndentedString(data)).append("\n");
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

