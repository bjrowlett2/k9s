// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/derailed/k9s/resource (interfaces: SwitchableRes)

package resource_test

import (
	k8s "github.com/derailed/k9s/resource/k8s"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockSwitchableRes struct {
	fail func(message string, callerSkip ...int)
}

func NewMockSwitchableRes() *MockSwitchableRes {
	return &MockSwitchableRes{fail: pegomock.GlobalFailHandler}
}

func (mock *MockSwitchableRes) Delete(_param0 string, _param1 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSwitchableRes().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Delete", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSwitchableRes) Get(_param0 string, _param1 string) (interface{}, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSwitchableRes().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Get", params, []reflect.Type{reflect.TypeOf((*interface{})(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 interface{}
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(interface{})
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockSwitchableRes) List(_param0 string) (k8s.Collection, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSwitchableRes().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("List", params, []reflect.Type{reflect.TypeOf((*k8s.Collection)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 k8s.Collection
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(k8s.Collection)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockSwitchableRes) Switch(_param0 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSwitchableRes().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Switch", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSwitchableRes) VerifyWasCalledOnce() *VerifierSwitchableRes {
	return &VerifierSwitchableRes{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockSwitchableRes) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierSwitchableRes {
	return &VerifierSwitchableRes{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockSwitchableRes) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierSwitchableRes {
	return &VerifierSwitchableRes{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockSwitchableRes) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierSwitchableRes {
	return &VerifierSwitchableRes{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierSwitchableRes struct {
	mock                   *MockSwitchableRes
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierSwitchableRes) Delete(_param0 string, _param1 string) *SwitchableRes_Delete_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Delete", params, verifier.timeout)
	return &SwitchableRes_Delete_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type SwitchableRes_Delete_OngoingVerification struct {
	mock              *MockSwitchableRes
	methodInvocations []pegomock.MethodInvocation
}

func (c *SwitchableRes_Delete_OngoingVerification) GetCapturedArguments() (string, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *SwitchableRes_Delete_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierSwitchableRes) Get(_param0 string, _param1 string) *SwitchableRes_Get_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Get", params, verifier.timeout)
	return &SwitchableRes_Get_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type SwitchableRes_Get_OngoingVerification struct {
	mock              *MockSwitchableRes
	methodInvocations []pegomock.MethodInvocation
}

func (c *SwitchableRes_Get_OngoingVerification) GetCapturedArguments() (string, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *SwitchableRes_Get_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierSwitchableRes) List(_param0 string) *SwitchableRes_List_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "List", params, verifier.timeout)
	return &SwitchableRes_List_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type SwitchableRes_List_OngoingVerification struct {
	mock              *MockSwitchableRes
	methodInvocations []pegomock.MethodInvocation
}

func (c *SwitchableRes_List_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *SwitchableRes_List_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierSwitchableRes) Switch(_param0 string) *SwitchableRes_Switch_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Switch", params, verifier.timeout)
	return &SwitchableRes_Switch_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type SwitchableRes_Switch_OngoingVerification struct {
	mock              *MockSwitchableRes
	methodInvocations []pegomock.MethodInvocation
}

func (c *SwitchableRes_Switch_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *SwitchableRes_Switch_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}
