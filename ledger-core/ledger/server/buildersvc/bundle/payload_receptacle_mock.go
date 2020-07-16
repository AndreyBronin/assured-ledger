package bundle

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/assured-ledger/ledger-core/vanilla/longbits"
)

// PayloadReceptacleMock implements PayloadReceptacle
type PayloadReceptacleMock struct {
	t minimock.Tester

	funcApplyFixedReader          func(f1 longbits.FixedReader) (err error)
	inspectFuncApplyFixedReader   func(f1 longbits.FixedReader)
	afterApplyFixedReaderCounter  uint64
	beforeApplyFixedReaderCounter uint64
	ApplyFixedReaderMock          mPayloadReceptacleMockApplyFixedReader

	funcApplyMarshalTo          func(m1 MarshalerTo) (err error)
	inspectFuncApplyMarshalTo   func(m1 MarshalerTo)
	afterApplyMarshalToCounter  uint64
	beforeApplyMarshalToCounter uint64
	ApplyMarshalToMock          mPayloadReceptacleMockApplyMarshalTo
}

// NewPayloadReceptacleMock returns a mock for PayloadReceptacle
func NewPayloadReceptacleMock(t minimock.Tester) *PayloadReceptacleMock {
	m := &PayloadReceptacleMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ApplyFixedReaderMock = mPayloadReceptacleMockApplyFixedReader{mock: m}
	m.ApplyFixedReaderMock.callArgs = []*PayloadReceptacleMockApplyFixedReaderParams{}

	m.ApplyMarshalToMock = mPayloadReceptacleMockApplyMarshalTo{mock: m}
	m.ApplyMarshalToMock.callArgs = []*PayloadReceptacleMockApplyMarshalToParams{}

	return m
}

type mPayloadReceptacleMockApplyFixedReader struct {
	mock               *PayloadReceptacleMock
	defaultExpectation *PayloadReceptacleMockApplyFixedReaderExpectation
	expectations       []*PayloadReceptacleMockApplyFixedReaderExpectation

	callArgs []*PayloadReceptacleMockApplyFixedReaderParams
	mutex    sync.RWMutex
}

// PayloadReceptacleMockApplyFixedReaderExpectation specifies expectation struct of the PayloadReceptacle.ApplyFixedReader
type PayloadReceptacleMockApplyFixedReaderExpectation struct {
	mock    *PayloadReceptacleMock
	params  *PayloadReceptacleMockApplyFixedReaderParams
	results *PayloadReceptacleMockApplyFixedReaderResults
	Counter uint64
}

// PayloadReceptacleMockApplyFixedReaderParams contains parameters of the PayloadReceptacle.ApplyFixedReader
type PayloadReceptacleMockApplyFixedReaderParams struct {
	f1 longbits.FixedReader
}

// PayloadReceptacleMockApplyFixedReaderResults contains results of the PayloadReceptacle.ApplyFixedReader
type PayloadReceptacleMockApplyFixedReaderResults struct {
	err error
}

// Expect sets up expected params for PayloadReceptacle.ApplyFixedReader
func (mmApplyFixedReader *mPayloadReceptacleMockApplyFixedReader) Expect(f1 longbits.FixedReader) *mPayloadReceptacleMockApplyFixedReader {
	if mmApplyFixedReader.mock.funcApplyFixedReader != nil {
		mmApplyFixedReader.mock.t.Fatalf("PayloadReceptacleMock.ApplyFixedReader mock is already set by Set")
	}

	if mmApplyFixedReader.defaultExpectation == nil {
		mmApplyFixedReader.defaultExpectation = &PayloadReceptacleMockApplyFixedReaderExpectation{}
	}

	mmApplyFixedReader.defaultExpectation.params = &PayloadReceptacleMockApplyFixedReaderParams{f1}
	for _, e := range mmApplyFixedReader.expectations {
		if minimock.Equal(e.params, mmApplyFixedReader.defaultExpectation.params) {
			mmApplyFixedReader.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmApplyFixedReader.defaultExpectation.params)
		}
	}

	return mmApplyFixedReader
}

// Inspect accepts an inspector function that has same arguments as the PayloadReceptacle.ApplyFixedReader
func (mmApplyFixedReader *mPayloadReceptacleMockApplyFixedReader) Inspect(f func(f1 longbits.FixedReader)) *mPayloadReceptacleMockApplyFixedReader {
	if mmApplyFixedReader.mock.inspectFuncApplyFixedReader != nil {
		mmApplyFixedReader.mock.t.Fatalf("Inspect function is already set for PayloadReceptacleMock.ApplyFixedReader")
	}

	mmApplyFixedReader.mock.inspectFuncApplyFixedReader = f

	return mmApplyFixedReader
}

// Return sets up results that will be returned by PayloadReceptacle.ApplyFixedReader
func (mmApplyFixedReader *mPayloadReceptacleMockApplyFixedReader) Return(err error) *PayloadReceptacleMock {
	if mmApplyFixedReader.mock.funcApplyFixedReader != nil {
		mmApplyFixedReader.mock.t.Fatalf("PayloadReceptacleMock.ApplyFixedReader mock is already set by Set")
	}

	if mmApplyFixedReader.defaultExpectation == nil {
		mmApplyFixedReader.defaultExpectation = &PayloadReceptacleMockApplyFixedReaderExpectation{mock: mmApplyFixedReader.mock}
	}
	mmApplyFixedReader.defaultExpectation.results = &PayloadReceptacleMockApplyFixedReaderResults{err}
	return mmApplyFixedReader.mock
}

//Set uses given function f to mock the PayloadReceptacle.ApplyFixedReader method
func (mmApplyFixedReader *mPayloadReceptacleMockApplyFixedReader) Set(f func(f1 longbits.FixedReader) (err error)) *PayloadReceptacleMock {
	if mmApplyFixedReader.defaultExpectation != nil {
		mmApplyFixedReader.mock.t.Fatalf("Default expectation is already set for the PayloadReceptacle.ApplyFixedReader method")
	}

	if len(mmApplyFixedReader.expectations) > 0 {
		mmApplyFixedReader.mock.t.Fatalf("Some expectations are already set for the PayloadReceptacle.ApplyFixedReader method")
	}

	mmApplyFixedReader.mock.funcApplyFixedReader = f
	return mmApplyFixedReader.mock
}

// When sets expectation for the PayloadReceptacle.ApplyFixedReader which will trigger the result defined by the following
// Then helper
func (mmApplyFixedReader *mPayloadReceptacleMockApplyFixedReader) When(f1 longbits.FixedReader) *PayloadReceptacleMockApplyFixedReaderExpectation {
	if mmApplyFixedReader.mock.funcApplyFixedReader != nil {
		mmApplyFixedReader.mock.t.Fatalf("PayloadReceptacleMock.ApplyFixedReader mock is already set by Set")
	}

	expectation := &PayloadReceptacleMockApplyFixedReaderExpectation{
		mock:   mmApplyFixedReader.mock,
		params: &PayloadReceptacleMockApplyFixedReaderParams{f1},
	}
	mmApplyFixedReader.expectations = append(mmApplyFixedReader.expectations, expectation)
	return expectation
}

// Then sets up PayloadReceptacle.ApplyFixedReader return parameters for the expectation previously defined by the When method
func (e *PayloadReceptacleMockApplyFixedReaderExpectation) Then(err error) *PayloadReceptacleMock {
	e.results = &PayloadReceptacleMockApplyFixedReaderResults{err}
	return e.mock
}

// ApplyFixedReader implements PayloadReceptacle
func (mmApplyFixedReader *PayloadReceptacleMock) ApplyFixedReader(f1 longbits.FixedReader) (err error) {
	mm_atomic.AddUint64(&mmApplyFixedReader.beforeApplyFixedReaderCounter, 1)
	defer mm_atomic.AddUint64(&mmApplyFixedReader.afterApplyFixedReaderCounter, 1)

	if mmApplyFixedReader.inspectFuncApplyFixedReader != nil {
		mmApplyFixedReader.inspectFuncApplyFixedReader(f1)
	}

	mm_params := &PayloadReceptacleMockApplyFixedReaderParams{f1}

	// Record call args
	mmApplyFixedReader.ApplyFixedReaderMock.mutex.Lock()
	mmApplyFixedReader.ApplyFixedReaderMock.callArgs = append(mmApplyFixedReader.ApplyFixedReaderMock.callArgs, mm_params)
	mmApplyFixedReader.ApplyFixedReaderMock.mutex.Unlock()

	for _, e := range mmApplyFixedReader.ApplyFixedReaderMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmApplyFixedReader.ApplyFixedReaderMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmApplyFixedReader.ApplyFixedReaderMock.defaultExpectation.Counter, 1)
		mm_want := mmApplyFixedReader.ApplyFixedReaderMock.defaultExpectation.params
		mm_got := PayloadReceptacleMockApplyFixedReaderParams{f1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmApplyFixedReader.t.Errorf("PayloadReceptacleMock.ApplyFixedReader got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmApplyFixedReader.ApplyFixedReaderMock.defaultExpectation.results
		if mm_results == nil {
			mmApplyFixedReader.t.Fatal("No results are set for the PayloadReceptacleMock.ApplyFixedReader")
		}
		return (*mm_results).err
	}
	if mmApplyFixedReader.funcApplyFixedReader != nil {
		return mmApplyFixedReader.funcApplyFixedReader(f1)
	}
	mmApplyFixedReader.t.Fatalf("Unexpected call to PayloadReceptacleMock.ApplyFixedReader. %v", f1)
	return
}

// ApplyFixedReaderAfterCounter returns a count of finished PayloadReceptacleMock.ApplyFixedReader invocations
func (mmApplyFixedReader *PayloadReceptacleMock) ApplyFixedReaderAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmApplyFixedReader.afterApplyFixedReaderCounter)
}

// ApplyFixedReaderBeforeCounter returns a count of PayloadReceptacleMock.ApplyFixedReader invocations
func (mmApplyFixedReader *PayloadReceptacleMock) ApplyFixedReaderBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmApplyFixedReader.beforeApplyFixedReaderCounter)
}

// Calls returns a list of arguments used in each call to PayloadReceptacleMock.ApplyFixedReader.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmApplyFixedReader *mPayloadReceptacleMockApplyFixedReader) Calls() []*PayloadReceptacleMockApplyFixedReaderParams {
	mmApplyFixedReader.mutex.RLock()

	argCopy := make([]*PayloadReceptacleMockApplyFixedReaderParams, len(mmApplyFixedReader.callArgs))
	copy(argCopy, mmApplyFixedReader.callArgs)

	mmApplyFixedReader.mutex.RUnlock()

	return argCopy
}

// MinimockApplyFixedReaderDone returns true if the count of the ApplyFixedReader invocations corresponds
// the number of defined expectations
func (m *PayloadReceptacleMock) MinimockApplyFixedReaderDone() bool {
	for _, e := range m.ApplyFixedReaderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ApplyFixedReaderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterApplyFixedReaderCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcApplyFixedReader != nil && mm_atomic.LoadUint64(&m.afterApplyFixedReaderCounter) < 1 {
		return false
	}
	return true
}

// MinimockApplyFixedReaderInspect logs each unmet expectation
func (m *PayloadReceptacleMock) MinimockApplyFixedReaderInspect() {
	for _, e := range m.ApplyFixedReaderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PayloadReceptacleMock.ApplyFixedReader with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ApplyFixedReaderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterApplyFixedReaderCounter) < 1 {
		if m.ApplyFixedReaderMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PayloadReceptacleMock.ApplyFixedReader")
		} else {
			m.t.Errorf("Expected call to PayloadReceptacleMock.ApplyFixedReader with params: %#v", *m.ApplyFixedReaderMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcApplyFixedReader != nil && mm_atomic.LoadUint64(&m.afterApplyFixedReaderCounter) < 1 {
		m.t.Error("Expected call to PayloadReceptacleMock.ApplyFixedReader")
	}
}

type mPayloadReceptacleMockApplyMarshalTo struct {
	mock               *PayloadReceptacleMock
	defaultExpectation *PayloadReceptacleMockApplyMarshalToExpectation
	expectations       []*PayloadReceptacleMockApplyMarshalToExpectation

	callArgs []*PayloadReceptacleMockApplyMarshalToParams
	mutex    sync.RWMutex
}

// PayloadReceptacleMockApplyMarshalToExpectation specifies expectation struct of the PayloadReceptacle.ApplyMarshalTo
type PayloadReceptacleMockApplyMarshalToExpectation struct {
	mock    *PayloadReceptacleMock
	params  *PayloadReceptacleMockApplyMarshalToParams
	results *PayloadReceptacleMockApplyMarshalToResults
	Counter uint64
}

// PayloadReceptacleMockApplyMarshalToParams contains parameters of the PayloadReceptacle.ApplyMarshalTo
type PayloadReceptacleMockApplyMarshalToParams struct {
	m1 MarshalerTo
}

// PayloadReceptacleMockApplyMarshalToResults contains results of the PayloadReceptacle.ApplyMarshalTo
type PayloadReceptacleMockApplyMarshalToResults struct {
	err error
}

// Expect sets up expected params for PayloadReceptacle.ApplyMarshalTo
func (mmApplyMarshalTo *mPayloadReceptacleMockApplyMarshalTo) Expect(m1 MarshalerTo) *mPayloadReceptacleMockApplyMarshalTo {
	if mmApplyMarshalTo.mock.funcApplyMarshalTo != nil {
		mmApplyMarshalTo.mock.t.Fatalf("PayloadReceptacleMock.ApplyMarshalTo mock is already set by Set")
	}

	if mmApplyMarshalTo.defaultExpectation == nil {
		mmApplyMarshalTo.defaultExpectation = &PayloadReceptacleMockApplyMarshalToExpectation{}
	}

	mmApplyMarshalTo.defaultExpectation.params = &PayloadReceptacleMockApplyMarshalToParams{m1}
	for _, e := range mmApplyMarshalTo.expectations {
		if minimock.Equal(e.params, mmApplyMarshalTo.defaultExpectation.params) {
			mmApplyMarshalTo.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmApplyMarshalTo.defaultExpectation.params)
		}
	}

	return mmApplyMarshalTo
}

// Inspect accepts an inspector function that has same arguments as the PayloadReceptacle.ApplyMarshalTo
func (mmApplyMarshalTo *mPayloadReceptacleMockApplyMarshalTo) Inspect(f func(m1 MarshalerTo)) *mPayloadReceptacleMockApplyMarshalTo {
	if mmApplyMarshalTo.mock.inspectFuncApplyMarshalTo != nil {
		mmApplyMarshalTo.mock.t.Fatalf("Inspect function is already set for PayloadReceptacleMock.ApplyMarshalTo")
	}

	mmApplyMarshalTo.mock.inspectFuncApplyMarshalTo = f

	return mmApplyMarshalTo
}

// Return sets up results that will be returned by PayloadReceptacle.ApplyMarshalTo
func (mmApplyMarshalTo *mPayloadReceptacleMockApplyMarshalTo) Return(err error) *PayloadReceptacleMock {
	if mmApplyMarshalTo.mock.funcApplyMarshalTo != nil {
		mmApplyMarshalTo.mock.t.Fatalf("PayloadReceptacleMock.ApplyMarshalTo mock is already set by Set")
	}

	if mmApplyMarshalTo.defaultExpectation == nil {
		mmApplyMarshalTo.defaultExpectation = &PayloadReceptacleMockApplyMarshalToExpectation{mock: mmApplyMarshalTo.mock}
	}
	mmApplyMarshalTo.defaultExpectation.results = &PayloadReceptacleMockApplyMarshalToResults{err}
	return mmApplyMarshalTo.mock
}

//Set uses given function f to mock the PayloadReceptacle.ApplyMarshalTo method
func (mmApplyMarshalTo *mPayloadReceptacleMockApplyMarshalTo) Set(f func(m1 MarshalerTo) (err error)) *PayloadReceptacleMock {
	if mmApplyMarshalTo.defaultExpectation != nil {
		mmApplyMarshalTo.mock.t.Fatalf("Default expectation is already set for the PayloadReceptacle.ApplyMarshalTo method")
	}

	if len(mmApplyMarshalTo.expectations) > 0 {
		mmApplyMarshalTo.mock.t.Fatalf("Some expectations are already set for the PayloadReceptacle.ApplyMarshalTo method")
	}

	mmApplyMarshalTo.mock.funcApplyMarshalTo = f
	return mmApplyMarshalTo.mock
}

// When sets expectation for the PayloadReceptacle.ApplyMarshalTo which will trigger the result defined by the following
// Then helper
func (mmApplyMarshalTo *mPayloadReceptacleMockApplyMarshalTo) When(m1 MarshalerTo) *PayloadReceptacleMockApplyMarshalToExpectation {
	if mmApplyMarshalTo.mock.funcApplyMarshalTo != nil {
		mmApplyMarshalTo.mock.t.Fatalf("PayloadReceptacleMock.ApplyMarshalTo mock is already set by Set")
	}

	expectation := &PayloadReceptacleMockApplyMarshalToExpectation{
		mock:   mmApplyMarshalTo.mock,
		params: &PayloadReceptacleMockApplyMarshalToParams{m1},
	}
	mmApplyMarshalTo.expectations = append(mmApplyMarshalTo.expectations, expectation)
	return expectation
}

// Then sets up PayloadReceptacle.ApplyMarshalTo return parameters for the expectation previously defined by the When method
func (e *PayloadReceptacleMockApplyMarshalToExpectation) Then(err error) *PayloadReceptacleMock {
	e.results = &PayloadReceptacleMockApplyMarshalToResults{err}
	return e.mock
}

// ApplyMarshalTo implements PayloadReceptacle
func (mmApplyMarshalTo *PayloadReceptacleMock) ApplyMarshalTo(m1 MarshalerTo) (err error) {
	mm_atomic.AddUint64(&mmApplyMarshalTo.beforeApplyMarshalToCounter, 1)
	defer mm_atomic.AddUint64(&mmApplyMarshalTo.afterApplyMarshalToCounter, 1)

	if mmApplyMarshalTo.inspectFuncApplyMarshalTo != nil {
		mmApplyMarshalTo.inspectFuncApplyMarshalTo(m1)
	}

	mm_params := &PayloadReceptacleMockApplyMarshalToParams{m1}

	// Record call args
	mmApplyMarshalTo.ApplyMarshalToMock.mutex.Lock()
	mmApplyMarshalTo.ApplyMarshalToMock.callArgs = append(mmApplyMarshalTo.ApplyMarshalToMock.callArgs, mm_params)
	mmApplyMarshalTo.ApplyMarshalToMock.mutex.Unlock()

	for _, e := range mmApplyMarshalTo.ApplyMarshalToMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmApplyMarshalTo.ApplyMarshalToMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmApplyMarshalTo.ApplyMarshalToMock.defaultExpectation.Counter, 1)
		mm_want := mmApplyMarshalTo.ApplyMarshalToMock.defaultExpectation.params
		mm_got := PayloadReceptacleMockApplyMarshalToParams{m1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmApplyMarshalTo.t.Errorf("PayloadReceptacleMock.ApplyMarshalTo got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmApplyMarshalTo.ApplyMarshalToMock.defaultExpectation.results
		if mm_results == nil {
			mmApplyMarshalTo.t.Fatal("No results are set for the PayloadReceptacleMock.ApplyMarshalTo")
		}
		return (*mm_results).err
	}
	if mmApplyMarshalTo.funcApplyMarshalTo != nil {
		return mmApplyMarshalTo.funcApplyMarshalTo(m1)
	}
	mmApplyMarshalTo.t.Fatalf("Unexpected call to PayloadReceptacleMock.ApplyMarshalTo. %v", m1)
	return
}

// ApplyMarshalToAfterCounter returns a count of finished PayloadReceptacleMock.ApplyMarshalTo invocations
func (mmApplyMarshalTo *PayloadReceptacleMock) ApplyMarshalToAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmApplyMarshalTo.afterApplyMarshalToCounter)
}

// ApplyMarshalToBeforeCounter returns a count of PayloadReceptacleMock.ApplyMarshalTo invocations
func (mmApplyMarshalTo *PayloadReceptacleMock) ApplyMarshalToBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmApplyMarshalTo.beforeApplyMarshalToCounter)
}

// Calls returns a list of arguments used in each call to PayloadReceptacleMock.ApplyMarshalTo.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmApplyMarshalTo *mPayloadReceptacleMockApplyMarshalTo) Calls() []*PayloadReceptacleMockApplyMarshalToParams {
	mmApplyMarshalTo.mutex.RLock()

	argCopy := make([]*PayloadReceptacleMockApplyMarshalToParams, len(mmApplyMarshalTo.callArgs))
	copy(argCopy, mmApplyMarshalTo.callArgs)

	mmApplyMarshalTo.mutex.RUnlock()

	return argCopy
}

// MinimockApplyMarshalToDone returns true if the count of the ApplyMarshalTo invocations corresponds
// the number of defined expectations
func (m *PayloadReceptacleMock) MinimockApplyMarshalToDone() bool {
	for _, e := range m.ApplyMarshalToMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ApplyMarshalToMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterApplyMarshalToCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcApplyMarshalTo != nil && mm_atomic.LoadUint64(&m.afterApplyMarshalToCounter) < 1 {
		return false
	}
	return true
}

// MinimockApplyMarshalToInspect logs each unmet expectation
func (m *PayloadReceptacleMock) MinimockApplyMarshalToInspect() {
	for _, e := range m.ApplyMarshalToMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PayloadReceptacleMock.ApplyMarshalTo with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ApplyMarshalToMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterApplyMarshalToCounter) < 1 {
		if m.ApplyMarshalToMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PayloadReceptacleMock.ApplyMarshalTo")
		} else {
			m.t.Errorf("Expected call to PayloadReceptacleMock.ApplyMarshalTo with params: %#v", *m.ApplyMarshalToMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcApplyMarshalTo != nil && mm_atomic.LoadUint64(&m.afterApplyMarshalToCounter) < 1 {
		m.t.Error("Expected call to PayloadReceptacleMock.ApplyMarshalTo")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PayloadReceptacleMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockApplyFixedReaderInspect()

		m.MinimockApplyMarshalToInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PayloadReceptacleMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *PayloadReceptacleMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockApplyFixedReaderDone() &&
		m.MinimockApplyMarshalToDone()
}