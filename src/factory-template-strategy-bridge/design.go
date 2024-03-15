package main

import (
	"errors"
	"fmt"
)

type InsureTemplateI interface {
	ExecuteInsure() (*Response, error)
	InsureHandlerI
}

type InsureHandlerI interface {
	ValidateRequest() error

	CheckRepetition() error
	// FillInClientRequest decrypt specific data into request's specific field
	FillInClientRequest() error
	// PrepareOrderCreateParam Prepare and get basic data for order creation
	PrepareOrderCreateParam() error
	LoadInsurePayStrategy() error
	// HandleInsure Handle the process of create order
	HandleInsure() (*Response, error)
}


type InsurePayStrategyI interface {
  
	DoPayCreateLogic(*OrderCreateParam, *InsuringStore) error

	GetInsureCreateResponse() *Response
}


// InsureTemplate working as an abstract class
// it only implements the template method `ExecuteInsure`
// In this `ExecuteInsure` method we will define the exact
// steps for this insure business.
type InsureTemplate struct {
	InsureHandlerI
}

func (template InsureTemplate) ExecuteInsure() (*Response, error) {
	fmt.Println("ValidateRequest")
	err := template.InsureHandlerI.ValidateRequest()
	if err != nil {
		return nil, err
	}
	err = template.InsureHandlerI.CheckRepetition()
	if err != nil {
		return nil, err
	}
	err = template.InsureHandlerI.PrepareOrderCreateParam()
	if err != nil {
		return nil, err
	}
	fmt.Println("LoadInsurePayStrategy")
	err = template.InsureHandlerI.LoadInsurePayStrategy()
	if err != nil {
		return nil, err
	}
	response, err := template.HandleInsure()
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ------------- InsureTemplate ends ------------- //



// GeneralInsureHandler working as a base class for all insure handler,
// and it implements the common implementation for handle insure.
// If other more specific handler have their own logic to implement these methods
// then over right them.
// Since go is incapable of having hierarchy like oop language,
// so we use composition that compose this `GeneralInsureHandler` into specific handler
// to overwrite the common version of InsureHandlerI's methods.
type GeneralInsureHandler struct {
	// pay strategy
	Strategy InsurePayStrategyI
	//originalReqByte []byte

	// necessary data for creating data
	InsureCreateOrderReq *OrderCreateParam
	// client request
	ClientInsureReq *Request
	InsuringStore *InsuringStore
}

func (handler *GeneralInsureHandler) ValidateRequest() error {
	fmt.Println("in ValidateRequest")
	if handler.ClientInsureReq.FieldA == "" {
		return errors.New("lack of filed A")
	}
	return nil
}

func (handler *GeneralInsureHandler) CheckRepetition() error {
	if handler.ClientInsureReq.FieldB == "" {
		return errors.New("repetition happened")
	}
	return nil
}

func (handler *GeneralInsureHandler) FillInClientRequest() error {
	return nil
}

func (handler *GeneralInsureHandler) PrepareOrderCreateParam() error {
	handler.InsureCreateOrderReq = nil
	return nil
}

func (handler *GeneralInsureHandler) LoadInsurePayStrategy() error {
	fmt.Println("in LoadInsurePayStrategy")
	//handler.Strategy = new(strategy.InsureStrategyA)
	return nil
}

func (handler *GeneralInsureHandler) HandleInsure() (*Response, error) {
	err := handler.Strategy.DoPayCreateLogic(handler.InsureCreateOrderReq, handler.InsuringStore)
	if err != nil {
		return nil, err
	}
	reply := handler.Strategy.GetInsureCreateResponse()

	fmt.Println("HandleInsure success end")
	return reply, nil
}

// ------------- GeneralInsureHandler ends ------------- //


// SpecificInsureHandler represents concretes for InsureHandlerI
// that actual logics were written into.
type SpecificInsureHandler struct {
	GeneralInsureHandler
}

func (handler *SpecificInsureHandler) FillInClientRequest() error {
	return nil
}

func (handler *SpecificInsureHandler) LoadInsurePayStrategy() error {

	fmt.Println("Load strategy in SpecificInsureHandler")
	if handler.ClientInsureReq.FieldA == "a" {
		handler.Strategy = new(InsureStrategyA)
		return nil
	}

	return errors.New("load strategy failed")
}

// ------------- SpecificInsureHandler ends ------------- //



type InsureStrategyA struct {
	reply *Response
}

func (a *InsureStrategyA) DoPayCreateLogic(param *OrderCreateParam, insuringStore *InsuringStore) error {
	return nil
}

func (a *InsureStrategyA) GetInsureCreateResponse() *Response {
	// In this method we can intercept the reply and some customize data
	// based on the strategy's logic (If it does have the needs to customize them)
	// otherwise we will simply return the reply
	return a.reply
}



// Response represent the response of request execution
type Response struct {

}

type Request struct {
	FieldA string
	FieldB string
}

type OrderCreateParam struct {

}

// InsuringStore store data generated that inside business login
type InsuringStore struct {

}

// NewInsureTemplate a simple factory for manufacturing insure template
func NewInsureTemplate(clientRequest *Request, symbol string) *InsureTemplate {
	insureTemplate := new(InsureTemplate)
	switch symbol {
	case "a":
		insureHandler := new(SpecificInsureHandler)
		insureHandler.ClientInsureReq = clientRequest
		insureTemplate.InsureHandlerI = insureHandler

	}

	return insureTemplate
}

func main() {
	reqeust := &Request{
		FieldA: "a",
		FieldB: "b",
	}
	template := NewInsureTemplate(reqeust, "a")
	fmt.Println("start")
	template.ExecuteInsure()
}
