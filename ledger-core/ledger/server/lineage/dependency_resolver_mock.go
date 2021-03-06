package lineage

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/assured-ledger/ledger-core/reference"
)

// DependencyResolverMock implements DependencyResolver
type DependencyResolverMock struct {
	t minimock.Tester

	funcFindLineDependency          func(root reference.Holder, ref reference.LocalHolder) (r1 ResolvedDependency, err error)
	inspectFuncFindLineDependency   func(root reference.Holder, ref reference.LocalHolder)
	afterFindLineDependencyCounter  uint64
	beforeFindLineDependencyCounter uint64
	FindLineDependencyMock          mDependencyResolverMockFindLineDependency

	funcFindOtherDependency          func(ref reference.Holder) (r1 ResolvedDependency, err error)
	inspectFuncFindOtherDependency   func(ref reference.Holder)
	afterFindOtherDependencyCounter  uint64
	beforeFindOtherDependencyCounter uint64
	FindOtherDependencyMock          mDependencyResolverMockFindOtherDependency
}

// NewDependencyResolverMock returns a mock for DependencyResolver
func NewDependencyResolverMock(t minimock.Tester) *DependencyResolverMock {
	m := &DependencyResolverMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.FindLineDependencyMock = mDependencyResolverMockFindLineDependency{mock: m}
	m.FindLineDependencyMock.callArgs = []*DependencyResolverMockFindLineDependencyParams{}

	m.FindOtherDependencyMock = mDependencyResolverMockFindOtherDependency{mock: m}
	m.FindOtherDependencyMock.callArgs = []*DependencyResolverMockFindOtherDependencyParams{}

	return m
}

type mDependencyResolverMockFindLineDependency struct {
	mock               *DependencyResolverMock
	defaultExpectation *DependencyResolverMockFindLineDependencyExpectation
	expectations       []*DependencyResolverMockFindLineDependencyExpectation

	callArgs []*DependencyResolverMockFindLineDependencyParams
	mutex    sync.RWMutex
}

// DependencyResolverMockFindLineDependencyExpectation specifies expectation struct of the DependencyResolver.FindLineDependency
type DependencyResolverMockFindLineDependencyExpectation struct {
	mock    *DependencyResolverMock
	params  *DependencyResolverMockFindLineDependencyParams
	results *DependencyResolverMockFindLineDependencyResults
	Counter uint64
}

// DependencyResolverMockFindLineDependencyParams contains parameters of the DependencyResolver.FindLineDependency
type DependencyResolverMockFindLineDependencyParams struct {
	root reference.Holder
	ref  reference.LocalHolder
}

// DependencyResolverMockFindLineDependencyResults contains results of the DependencyResolver.FindLineDependency
type DependencyResolverMockFindLineDependencyResults struct {
	r1  ResolvedDependency
	err error
}

// Expect sets up expected params for DependencyResolver.FindLineDependency
func (mmFindLineDependency *mDependencyResolverMockFindLineDependency) Expect(root reference.Holder, ref reference.LocalHolder) *mDependencyResolverMockFindLineDependency {
	if mmFindLineDependency.mock.funcFindLineDependency != nil {
		mmFindLineDependency.mock.t.Fatalf("DependencyResolverMock.FindLineDependency mock is already set by Set")
	}

	if mmFindLineDependency.defaultExpectation == nil {
		mmFindLineDependency.defaultExpectation = &DependencyResolverMockFindLineDependencyExpectation{}
	}

	mmFindLineDependency.defaultExpectation.params = &DependencyResolverMockFindLineDependencyParams{root, ref}
	for _, e := range mmFindLineDependency.expectations {
		if minimock.Equal(e.params, mmFindLineDependency.defaultExpectation.params) {
			mmFindLineDependency.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFindLineDependency.defaultExpectation.params)
		}
	}

	return mmFindLineDependency
}

// Inspect accepts an inspector function that has same arguments as the DependencyResolver.FindLineDependency
func (mmFindLineDependency *mDependencyResolverMockFindLineDependency) Inspect(f func(root reference.Holder, ref reference.LocalHolder)) *mDependencyResolverMockFindLineDependency {
	if mmFindLineDependency.mock.inspectFuncFindLineDependency != nil {
		mmFindLineDependency.mock.t.Fatalf("Inspect function is already set for DependencyResolverMock.FindLineDependency")
	}

	mmFindLineDependency.mock.inspectFuncFindLineDependency = f

	return mmFindLineDependency
}

// Return sets up results that will be returned by DependencyResolver.FindLineDependency
func (mmFindLineDependency *mDependencyResolverMockFindLineDependency) Return(r1 ResolvedDependency, err error) *DependencyResolverMock {
	if mmFindLineDependency.mock.funcFindLineDependency != nil {
		mmFindLineDependency.mock.t.Fatalf("DependencyResolverMock.FindLineDependency mock is already set by Set")
	}

	if mmFindLineDependency.defaultExpectation == nil {
		mmFindLineDependency.defaultExpectation = &DependencyResolverMockFindLineDependencyExpectation{mock: mmFindLineDependency.mock}
	}
	mmFindLineDependency.defaultExpectation.results = &DependencyResolverMockFindLineDependencyResults{r1, err}
	return mmFindLineDependency.mock
}

//Set uses given function f to mock the DependencyResolver.FindLineDependency method
func (mmFindLineDependency *mDependencyResolverMockFindLineDependency) Set(f func(root reference.Holder, ref reference.LocalHolder) (r1 ResolvedDependency, err error)) *DependencyResolverMock {
	if mmFindLineDependency.defaultExpectation != nil {
		mmFindLineDependency.mock.t.Fatalf("Default expectation is already set for the DependencyResolver.FindLineDependency method")
	}

	if len(mmFindLineDependency.expectations) > 0 {
		mmFindLineDependency.mock.t.Fatalf("Some expectations are already set for the DependencyResolver.FindLineDependency method")
	}

	mmFindLineDependency.mock.funcFindLineDependency = f
	return mmFindLineDependency.mock
}

// When sets expectation for the DependencyResolver.FindLineDependency which will trigger the result defined by the following
// Then helper
func (mmFindLineDependency *mDependencyResolverMockFindLineDependency) When(root reference.Holder, ref reference.LocalHolder) *DependencyResolverMockFindLineDependencyExpectation {
	if mmFindLineDependency.mock.funcFindLineDependency != nil {
		mmFindLineDependency.mock.t.Fatalf("DependencyResolverMock.FindLineDependency mock is already set by Set")
	}

	expectation := &DependencyResolverMockFindLineDependencyExpectation{
		mock:   mmFindLineDependency.mock,
		params: &DependencyResolverMockFindLineDependencyParams{root, ref},
	}
	mmFindLineDependency.expectations = append(mmFindLineDependency.expectations, expectation)
	return expectation
}

// Then sets up DependencyResolver.FindLineDependency return parameters for the expectation previously defined by the When method
func (e *DependencyResolverMockFindLineDependencyExpectation) Then(r1 ResolvedDependency, err error) *DependencyResolverMock {
	e.results = &DependencyResolverMockFindLineDependencyResults{r1, err}
	return e.mock
}

// FindLineDependency implements DependencyResolver
func (mmFindLineDependency *DependencyResolverMock) FindLineDependency(root reference.Holder, ref reference.LocalHolder) (r1 ResolvedDependency, err error) {
	mm_atomic.AddUint64(&mmFindLineDependency.beforeFindLineDependencyCounter, 1)
	defer mm_atomic.AddUint64(&mmFindLineDependency.afterFindLineDependencyCounter, 1)

	if mmFindLineDependency.inspectFuncFindLineDependency != nil {
		mmFindLineDependency.inspectFuncFindLineDependency(root, ref)
	}

	mm_params := &DependencyResolverMockFindLineDependencyParams{root, ref}

	// Record call args
	mmFindLineDependency.FindLineDependencyMock.mutex.Lock()
	mmFindLineDependency.FindLineDependencyMock.callArgs = append(mmFindLineDependency.FindLineDependencyMock.callArgs, mm_params)
	mmFindLineDependency.FindLineDependencyMock.mutex.Unlock()

	for _, e := range mmFindLineDependency.FindLineDependencyMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.err
		}
	}

	if mmFindLineDependency.FindLineDependencyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFindLineDependency.FindLineDependencyMock.defaultExpectation.Counter, 1)
		mm_want := mmFindLineDependency.FindLineDependencyMock.defaultExpectation.params
		mm_got := DependencyResolverMockFindLineDependencyParams{root, ref}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFindLineDependency.t.Errorf("DependencyResolverMock.FindLineDependency got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmFindLineDependency.FindLineDependencyMock.defaultExpectation.results
		if mm_results == nil {
			mmFindLineDependency.t.Fatal("No results are set for the DependencyResolverMock.FindLineDependency")
		}
		return (*mm_results).r1, (*mm_results).err
	}
	if mmFindLineDependency.funcFindLineDependency != nil {
		return mmFindLineDependency.funcFindLineDependency(root, ref)
	}
	mmFindLineDependency.t.Fatalf("Unexpected call to DependencyResolverMock.FindLineDependency. %v %v", root, ref)
	return
}

// FindLineDependencyAfterCounter returns a count of finished DependencyResolverMock.FindLineDependency invocations
func (mmFindLineDependency *DependencyResolverMock) FindLineDependencyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindLineDependency.afterFindLineDependencyCounter)
}

// FindLineDependencyBeforeCounter returns a count of DependencyResolverMock.FindLineDependency invocations
func (mmFindLineDependency *DependencyResolverMock) FindLineDependencyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindLineDependency.beforeFindLineDependencyCounter)
}

// Calls returns a list of arguments used in each call to DependencyResolverMock.FindLineDependency.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFindLineDependency *mDependencyResolverMockFindLineDependency) Calls() []*DependencyResolverMockFindLineDependencyParams {
	mmFindLineDependency.mutex.RLock()

	argCopy := make([]*DependencyResolverMockFindLineDependencyParams, len(mmFindLineDependency.callArgs))
	copy(argCopy, mmFindLineDependency.callArgs)

	mmFindLineDependency.mutex.RUnlock()

	return argCopy
}

// MinimockFindLineDependencyDone returns true if the count of the FindLineDependency invocations corresponds
// the number of defined expectations
func (m *DependencyResolverMock) MinimockFindLineDependencyDone() bool {
	for _, e := range m.FindLineDependencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindLineDependencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFindLineDependencyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindLineDependency != nil && mm_atomic.LoadUint64(&m.afterFindLineDependencyCounter) < 1 {
		return false
	}
	return true
}

// MinimockFindLineDependencyInspect logs each unmet expectation
func (m *DependencyResolverMock) MinimockFindLineDependencyInspect() {
	for _, e := range m.FindLineDependencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DependencyResolverMock.FindLineDependency with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindLineDependencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFindLineDependencyCounter) < 1 {
		if m.FindLineDependencyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DependencyResolverMock.FindLineDependency")
		} else {
			m.t.Errorf("Expected call to DependencyResolverMock.FindLineDependency with params: %#v", *m.FindLineDependencyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindLineDependency != nil && mm_atomic.LoadUint64(&m.afterFindLineDependencyCounter) < 1 {
		m.t.Error("Expected call to DependencyResolverMock.FindLineDependency")
	}
}

type mDependencyResolverMockFindOtherDependency struct {
	mock               *DependencyResolverMock
	defaultExpectation *DependencyResolverMockFindOtherDependencyExpectation
	expectations       []*DependencyResolverMockFindOtherDependencyExpectation

	callArgs []*DependencyResolverMockFindOtherDependencyParams
	mutex    sync.RWMutex
}

// DependencyResolverMockFindOtherDependencyExpectation specifies expectation struct of the DependencyResolver.FindOtherDependency
type DependencyResolverMockFindOtherDependencyExpectation struct {
	mock    *DependencyResolverMock
	params  *DependencyResolverMockFindOtherDependencyParams
	results *DependencyResolverMockFindOtherDependencyResults
	Counter uint64
}

// DependencyResolverMockFindOtherDependencyParams contains parameters of the DependencyResolver.FindOtherDependency
type DependencyResolverMockFindOtherDependencyParams struct {
	ref reference.Holder
}

// DependencyResolverMockFindOtherDependencyResults contains results of the DependencyResolver.FindOtherDependency
type DependencyResolverMockFindOtherDependencyResults struct {
	r1  ResolvedDependency
	err error
}

// Expect sets up expected params for DependencyResolver.FindOtherDependency
func (mmFindOtherDependency *mDependencyResolverMockFindOtherDependency) Expect(ref reference.Holder) *mDependencyResolverMockFindOtherDependency {
	if mmFindOtherDependency.mock.funcFindOtherDependency != nil {
		mmFindOtherDependency.mock.t.Fatalf("DependencyResolverMock.FindOtherDependency mock is already set by Set")
	}

	if mmFindOtherDependency.defaultExpectation == nil {
		mmFindOtherDependency.defaultExpectation = &DependencyResolverMockFindOtherDependencyExpectation{}
	}

	mmFindOtherDependency.defaultExpectation.params = &DependencyResolverMockFindOtherDependencyParams{ref}
	for _, e := range mmFindOtherDependency.expectations {
		if minimock.Equal(e.params, mmFindOtherDependency.defaultExpectation.params) {
			mmFindOtherDependency.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFindOtherDependency.defaultExpectation.params)
		}
	}

	return mmFindOtherDependency
}

// Inspect accepts an inspector function that has same arguments as the DependencyResolver.FindOtherDependency
func (mmFindOtherDependency *mDependencyResolverMockFindOtherDependency) Inspect(f func(ref reference.Holder)) *mDependencyResolverMockFindOtherDependency {
	if mmFindOtherDependency.mock.inspectFuncFindOtherDependency != nil {
		mmFindOtherDependency.mock.t.Fatalf("Inspect function is already set for DependencyResolverMock.FindOtherDependency")
	}

	mmFindOtherDependency.mock.inspectFuncFindOtherDependency = f

	return mmFindOtherDependency
}

// Return sets up results that will be returned by DependencyResolver.FindOtherDependency
func (mmFindOtherDependency *mDependencyResolverMockFindOtherDependency) Return(r1 ResolvedDependency, err error) *DependencyResolverMock {
	if mmFindOtherDependency.mock.funcFindOtherDependency != nil {
		mmFindOtherDependency.mock.t.Fatalf("DependencyResolverMock.FindOtherDependency mock is already set by Set")
	}

	if mmFindOtherDependency.defaultExpectation == nil {
		mmFindOtherDependency.defaultExpectation = &DependencyResolverMockFindOtherDependencyExpectation{mock: mmFindOtherDependency.mock}
	}
	mmFindOtherDependency.defaultExpectation.results = &DependencyResolverMockFindOtherDependencyResults{r1, err}
	return mmFindOtherDependency.mock
}

//Set uses given function f to mock the DependencyResolver.FindOtherDependency method
func (mmFindOtherDependency *mDependencyResolverMockFindOtherDependency) Set(f func(ref reference.Holder) (r1 ResolvedDependency, err error)) *DependencyResolverMock {
	if mmFindOtherDependency.defaultExpectation != nil {
		mmFindOtherDependency.mock.t.Fatalf("Default expectation is already set for the DependencyResolver.FindOtherDependency method")
	}

	if len(mmFindOtherDependency.expectations) > 0 {
		mmFindOtherDependency.mock.t.Fatalf("Some expectations are already set for the DependencyResolver.FindOtherDependency method")
	}

	mmFindOtherDependency.mock.funcFindOtherDependency = f
	return mmFindOtherDependency.mock
}

// When sets expectation for the DependencyResolver.FindOtherDependency which will trigger the result defined by the following
// Then helper
func (mmFindOtherDependency *mDependencyResolverMockFindOtherDependency) When(ref reference.Holder) *DependencyResolverMockFindOtherDependencyExpectation {
	if mmFindOtherDependency.mock.funcFindOtherDependency != nil {
		mmFindOtherDependency.mock.t.Fatalf("DependencyResolverMock.FindOtherDependency mock is already set by Set")
	}

	expectation := &DependencyResolverMockFindOtherDependencyExpectation{
		mock:   mmFindOtherDependency.mock,
		params: &DependencyResolverMockFindOtherDependencyParams{ref},
	}
	mmFindOtherDependency.expectations = append(mmFindOtherDependency.expectations, expectation)
	return expectation
}

// Then sets up DependencyResolver.FindOtherDependency return parameters for the expectation previously defined by the When method
func (e *DependencyResolverMockFindOtherDependencyExpectation) Then(r1 ResolvedDependency, err error) *DependencyResolverMock {
	e.results = &DependencyResolverMockFindOtherDependencyResults{r1, err}
	return e.mock
}

// FindOtherDependency implements DependencyResolver
func (mmFindOtherDependency *DependencyResolverMock) FindOtherDependency(ref reference.Holder) (r1 ResolvedDependency, err error) {
	mm_atomic.AddUint64(&mmFindOtherDependency.beforeFindOtherDependencyCounter, 1)
	defer mm_atomic.AddUint64(&mmFindOtherDependency.afterFindOtherDependencyCounter, 1)

	if mmFindOtherDependency.inspectFuncFindOtherDependency != nil {
		mmFindOtherDependency.inspectFuncFindOtherDependency(ref)
	}

	mm_params := &DependencyResolverMockFindOtherDependencyParams{ref}

	// Record call args
	mmFindOtherDependency.FindOtherDependencyMock.mutex.Lock()
	mmFindOtherDependency.FindOtherDependencyMock.callArgs = append(mmFindOtherDependency.FindOtherDependencyMock.callArgs, mm_params)
	mmFindOtherDependency.FindOtherDependencyMock.mutex.Unlock()

	for _, e := range mmFindOtherDependency.FindOtherDependencyMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.err
		}
	}

	if mmFindOtherDependency.FindOtherDependencyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFindOtherDependency.FindOtherDependencyMock.defaultExpectation.Counter, 1)
		mm_want := mmFindOtherDependency.FindOtherDependencyMock.defaultExpectation.params
		mm_got := DependencyResolverMockFindOtherDependencyParams{ref}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFindOtherDependency.t.Errorf("DependencyResolverMock.FindOtherDependency got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmFindOtherDependency.FindOtherDependencyMock.defaultExpectation.results
		if mm_results == nil {
			mmFindOtherDependency.t.Fatal("No results are set for the DependencyResolverMock.FindOtherDependency")
		}
		return (*mm_results).r1, (*mm_results).err
	}
	if mmFindOtherDependency.funcFindOtherDependency != nil {
		return mmFindOtherDependency.funcFindOtherDependency(ref)
	}
	mmFindOtherDependency.t.Fatalf("Unexpected call to DependencyResolverMock.FindOtherDependency. %v", ref)
	return
}

// FindOtherDependencyAfterCounter returns a count of finished DependencyResolverMock.FindOtherDependency invocations
func (mmFindOtherDependency *DependencyResolverMock) FindOtherDependencyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindOtherDependency.afterFindOtherDependencyCounter)
}

// FindOtherDependencyBeforeCounter returns a count of DependencyResolverMock.FindOtherDependency invocations
func (mmFindOtherDependency *DependencyResolverMock) FindOtherDependencyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindOtherDependency.beforeFindOtherDependencyCounter)
}

// Calls returns a list of arguments used in each call to DependencyResolverMock.FindOtherDependency.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFindOtherDependency *mDependencyResolverMockFindOtherDependency) Calls() []*DependencyResolverMockFindOtherDependencyParams {
	mmFindOtherDependency.mutex.RLock()

	argCopy := make([]*DependencyResolverMockFindOtherDependencyParams, len(mmFindOtherDependency.callArgs))
	copy(argCopy, mmFindOtherDependency.callArgs)

	mmFindOtherDependency.mutex.RUnlock()

	return argCopy
}

// MinimockFindOtherDependencyDone returns true if the count of the FindOtherDependency invocations corresponds
// the number of defined expectations
func (m *DependencyResolverMock) MinimockFindOtherDependencyDone() bool {
	for _, e := range m.FindOtherDependencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindOtherDependencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFindOtherDependencyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindOtherDependency != nil && mm_atomic.LoadUint64(&m.afterFindOtherDependencyCounter) < 1 {
		return false
	}
	return true
}

// MinimockFindOtherDependencyInspect logs each unmet expectation
func (m *DependencyResolverMock) MinimockFindOtherDependencyInspect() {
	for _, e := range m.FindOtherDependencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DependencyResolverMock.FindOtherDependency with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindOtherDependencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFindOtherDependencyCounter) < 1 {
		if m.FindOtherDependencyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DependencyResolverMock.FindOtherDependency")
		} else {
			m.t.Errorf("Expected call to DependencyResolverMock.FindOtherDependency with params: %#v", *m.FindOtherDependencyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindOtherDependency != nil && mm_atomic.LoadUint64(&m.afterFindOtherDependencyCounter) < 1 {
		m.t.Error("Expected call to DependencyResolverMock.FindOtherDependency")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DependencyResolverMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockFindLineDependencyInspect()

		m.MinimockFindOtherDependencyInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DependencyResolverMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *DependencyResolverMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockFindLineDependencyDone() &&
		m.MinimockFindOtherDependencyDone()
}
