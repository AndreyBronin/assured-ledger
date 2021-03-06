package profiles

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/assured-ledger/ledger-core/insolar/node"
	"github.com/insolar/assured-ledger/ledger-core/network/consensus/gcpv2/api/member"
	"github.com/insolar/assured-ledger/ledger-core/vanilla/cryptkit"
)

// BaseNodeMock implements BaseNode
type BaseNodeMock struct {
	t minimock.Tester

	funcGetNodeID          func() (s1 node.ShortNodeID)
	inspectFuncGetNodeID   func()
	afterGetNodeIDCounter  uint64
	beforeGetNodeIDCounter uint64
	GetNodeIDMock          mBaseNodeMockGetNodeID

	funcGetOpMode          func() (o1 member.OpMode)
	inspectFuncGetOpMode   func()
	afterGetOpModeCounter  uint64
	beforeGetOpModeCounter uint64
	GetOpModeMock          mBaseNodeMockGetOpMode

	funcGetSignatureVerifier          func() (s1 cryptkit.SignatureVerifier)
	inspectFuncGetSignatureVerifier   func()
	afterGetSignatureVerifierCounter  uint64
	beforeGetSignatureVerifierCounter uint64
	GetSignatureVerifierMock          mBaseNodeMockGetSignatureVerifier

	funcGetStatic          func() (s1 StaticProfile)
	inspectFuncGetStatic   func()
	afterGetStaticCounter  uint64
	beforeGetStaticCounter uint64
	GetStaticMock          mBaseNodeMockGetStatic
}

// NewBaseNodeMock returns a mock for BaseNode
func NewBaseNodeMock(t minimock.Tester) *BaseNodeMock {
	m := &BaseNodeMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetNodeIDMock = mBaseNodeMockGetNodeID{mock: m}

	m.GetOpModeMock = mBaseNodeMockGetOpMode{mock: m}

	m.GetSignatureVerifierMock = mBaseNodeMockGetSignatureVerifier{mock: m}

	m.GetStaticMock = mBaseNodeMockGetStatic{mock: m}

	return m
}

type mBaseNodeMockGetNodeID struct {
	mock               *BaseNodeMock
	defaultExpectation *BaseNodeMockGetNodeIDExpectation
	expectations       []*BaseNodeMockGetNodeIDExpectation
}

// BaseNodeMockGetNodeIDExpectation specifies expectation struct of the BaseNode.GetNodeID
type BaseNodeMockGetNodeIDExpectation struct {
	mock *BaseNodeMock

	results *BaseNodeMockGetNodeIDResults
	Counter uint64
}

// BaseNodeMockGetNodeIDResults contains results of the BaseNode.GetNodeID
type BaseNodeMockGetNodeIDResults struct {
	s1 node.ShortNodeID
}

// Expect sets up expected params for BaseNode.GetNodeID
func (mmGetNodeID *mBaseNodeMockGetNodeID) Expect() *mBaseNodeMockGetNodeID {
	if mmGetNodeID.mock.funcGetNodeID != nil {
		mmGetNodeID.mock.t.Fatalf("BaseNodeMock.GetNodeID mock is already set by Set")
	}

	if mmGetNodeID.defaultExpectation == nil {
		mmGetNodeID.defaultExpectation = &BaseNodeMockGetNodeIDExpectation{}
	}

	return mmGetNodeID
}

// Inspect accepts an inspector function that has same arguments as the BaseNode.GetNodeID
func (mmGetNodeID *mBaseNodeMockGetNodeID) Inspect(f func()) *mBaseNodeMockGetNodeID {
	if mmGetNodeID.mock.inspectFuncGetNodeID != nil {
		mmGetNodeID.mock.t.Fatalf("Inspect function is already set for BaseNodeMock.GetNodeID")
	}

	mmGetNodeID.mock.inspectFuncGetNodeID = f

	return mmGetNodeID
}

// Return sets up results that will be returned by BaseNode.GetNodeID
func (mmGetNodeID *mBaseNodeMockGetNodeID) Return(s1 node.ShortNodeID) *BaseNodeMock {
	if mmGetNodeID.mock.funcGetNodeID != nil {
		mmGetNodeID.mock.t.Fatalf("BaseNodeMock.GetNodeID mock is already set by Set")
	}

	if mmGetNodeID.defaultExpectation == nil {
		mmGetNodeID.defaultExpectation = &BaseNodeMockGetNodeIDExpectation{mock: mmGetNodeID.mock}
	}
	mmGetNodeID.defaultExpectation.results = &BaseNodeMockGetNodeIDResults{s1}
	return mmGetNodeID.mock
}

//Set uses given function f to mock the BaseNode.GetNodeID method
func (mmGetNodeID *mBaseNodeMockGetNodeID) Set(f func() (s1 node.ShortNodeID)) *BaseNodeMock {
	if mmGetNodeID.defaultExpectation != nil {
		mmGetNodeID.mock.t.Fatalf("Default expectation is already set for the BaseNode.GetNodeID method")
	}

	if len(mmGetNodeID.expectations) > 0 {
		mmGetNodeID.mock.t.Fatalf("Some expectations are already set for the BaseNode.GetNodeID method")
	}

	mmGetNodeID.mock.funcGetNodeID = f
	return mmGetNodeID.mock
}

// GetNodeID implements BaseNode
func (mmGetNodeID *BaseNodeMock) GetNodeID() (s1 node.ShortNodeID) {
	mm_atomic.AddUint64(&mmGetNodeID.beforeGetNodeIDCounter, 1)
	defer mm_atomic.AddUint64(&mmGetNodeID.afterGetNodeIDCounter, 1)

	if mmGetNodeID.inspectFuncGetNodeID != nil {
		mmGetNodeID.inspectFuncGetNodeID()
	}

	if mmGetNodeID.GetNodeIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetNodeID.GetNodeIDMock.defaultExpectation.Counter, 1)

		mm_results := mmGetNodeID.GetNodeIDMock.defaultExpectation.results
		if mm_results == nil {
			mmGetNodeID.t.Fatal("No results are set for the BaseNodeMock.GetNodeID")
		}
		return (*mm_results).s1
	}
	if mmGetNodeID.funcGetNodeID != nil {
		return mmGetNodeID.funcGetNodeID()
	}
	mmGetNodeID.t.Fatalf("Unexpected call to BaseNodeMock.GetNodeID.")
	return
}

// GetNodeIDAfterCounter returns a count of finished BaseNodeMock.GetNodeID invocations
func (mmGetNodeID *BaseNodeMock) GetNodeIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetNodeID.afterGetNodeIDCounter)
}

// GetNodeIDBeforeCounter returns a count of BaseNodeMock.GetNodeID invocations
func (mmGetNodeID *BaseNodeMock) GetNodeIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetNodeID.beforeGetNodeIDCounter)
}

// MinimockGetNodeIDDone returns true if the count of the GetNodeID invocations corresponds
// the number of defined expectations
func (m *BaseNodeMock) MinimockGetNodeIDDone() bool {
	for _, e := range m.GetNodeIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetNodeIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetNodeIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetNodeID != nil && mm_atomic.LoadUint64(&m.afterGetNodeIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetNodeIDInspect logs each unmet expectation
func (m *BaseNodeMock) MinimockGetNodeIDInspect() {
	for _, e := range m.GetNodeIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to BaseNodeMock.GetNodeID")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetNodeIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetNodeIDCounter) < 1 {
		m.t.Error("Expected call to BaseNodeMock.GetNodeID")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetNodeID != nil && mm_atomic.LoadUint64(&m.afterGetNodeIDCounter) < 1 {
		m.t.Error("Expected call to BaseNodeMock.GetNodeID")
	}
}

type mBaseNodeMockGetOpMode struct {
	mock               *BaseNodeMock
	defaultExpectation *BaseNodeMockGetOpModeExpectation
	expectations       []*BaseNodeMockGetOpModeExpectation
}

// BaseNodeMockGetOpModeExpectation specifies expectation struct of the BaseNode.GetOpMode
type BaseNodeMockGetOpModeExpectation struct {
	mock *BaseNodeMock

	results *BaseNodeMockGetOpModeResults
	Counter uint64
}

// BaseNodeMockGetOpModeResults contains results of the BaseNode.GetOpMode
type BaseNodeMockGetOpModeResults struct {
	o1 member.OpMode
}

// Expect sets up expected params for BaseNode.GetOpMode
func (mmGetOpMode *mBaseNodeMockGetOpMode) Expect() *mBaseNodeMockGetOpMode {
	if mmGetOpMode.mock.funcGetOpMode != nil {
		mmGetOpMode.mock.t.Fatalf("BaseNodeMock.GetOpMode mock is already set by Set")
	}

	if mmGetOpMode.defaultExpectation == nil {
		mmGetOpMode.defaultExpectation = &BaseNodeMockGetOpModeExpectation{}
	}

	return mmGetOpMode
}

// Inspect accepts an inspector function that has same arguments as the BaseNode.GetOpMode
func (mmGetOpMode *mBaseNodeMockGetOpMode) Inspect(f func()) *mBaseNodeMockGetOpMode {
	if mmGetOpMode.mock.inspectFuncGetOpMode != nil {
		mmGetOpMode.mock.t.Fatalf("Inspect function is already set for BaseNodeMock.GetOpMode")
	}

	mmGetOpMode.mock.inspectFuncGetOpMode = f

	return mmGetOpMode
}

// Return sets up results that will be returned by BaseNode.GetOpMode
func (mmGetOpMode *mBaseNodeMockGetOpMode) Return(o1 member.OpMode) *BaseNodeMock {
	if mmGetOpMode.mock.funcGetOpMode != nil {
		mmGetOpMode.mock.t.Fatalf("BaseNodeMock.GetOpMode mock is already set by Set")
	}

	if mmGetOpMode.defaultExpectation == nil {
		mmGetOpMode.defaultExpectation = &BaseNodeMockGetOpModeExpectation{mock: mmGetOpMode.mock}
	}
	mmGetOpMode.defaultExpectation.results = &BaseNodeMockGetOpModeResults{o1}
	return mmGetOpMode.mock
}

//Set uses given function f to mock the BaseNode.GetOpMode method
func (mmGetOpMode *mBaseNodeMockGetOpMode) Set(f func() (o1 member.OpMode)) *BaseNodeMock {
	if mmGetOpMode.defaultExpectation != nil {
		mmGetOpMode.mock.t.Fatalf("Default expectation is already set for the BaseNode.GetOpMode method")
	}

	if len(mmGetOpMode.expectations) > 0 {
		mmGetOpMode.mock.t.Fatalf("Some expectations are already set for the BaseNode.GetOpMode method")
	}

	mmGetOpMode.mock.funcGetOpMode = f
	return mmGetOpMode.mock
}

// GetOpMode implements BaseNode
func (mmGetOpMode *BaseNodeMock) GetOpMode() (o1 member.OpMode) {
	mm_atomic.AddUint64(&mmGetOpMode.beforeGetOpModeCounter, 1)
	defer mm_atomic.AddUint64(&mmGetOpMode.afterGetOpModeCounter, 1)

	if mmGetOpMode.inspectFuncGetOpMode != nil {
		mmGetOpMode.inspectFuncGetOpMode()
	}

	if mmGetOpMode.GetOpModeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetOpMode.GetOpModeMock.defaultExpectation.Counter, 1)

		mm_results := mmGetOpMode.GetOpModeMock.defaultExpectation.results
		if mm_results == nil {
			mmGetOpMode.t.Fatal("No results are set for the BaseNodeMock.GetOpMode")
		}
		return (*mm_results).o1
	}
	if mmGetOpMode.funcGetOpMode != nil {
		return mmGetOpMode.funcGetOpMode()
	}
	mmGetOpMode.t.Fatalf("Unexpected call to BaseNodeMock.GetOpMode.")
	return
}

// GetOpModeAfterCounter returns a count of finished BaseNodeMock.GetOpMode invocations
func (mmGetOpMode *BaseNodeMock) GetOpModeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetOpMode.afterGetOpModeCounter)
}

// GetOpModeBeforeCounter returns a count of BaseNodeMock.GetOpMode invocations
func (mmGetOpMode *BaseNodeMock) GetOpModeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetOpMode.beforeGetOpModeCounter)
}

// MinimockGetOpModeDone returns true if the count of the GetOpMode invocations corresponds
// the number of defined expectations
func (m *BaseNodeMock) MinimockGetOpModeDone() bool {
	for _, e := range m.GetOpModeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetOpModeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetOpModeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetOpMode != nil && mm_atomic.LoadUint64(&m.afterGetOpModeCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetOpModeInspect logs each unmet expectation
func (m *BaseNodeMock) MinimockGetOpModeInspect() {
	for _, e := range m.GetOpModeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to BaseNodeMock.GetOpMode")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetOpModeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetOpModeCounter) < 1 {
		m.t.Error("Expected call to BaseNodeMock.GetOpMode")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetOpMode != nil && mm_atomic.LoadUint64(&m.afterGetOpModeCounter) < 1 {
		m.t.Error("Expected call to BaseNodeMock.GetOpMode")
	}
}

type mBaseNodeMockGetSignatureVerifier struct {
	mock               *BaseNodeMock
	defaultExpectation *BaseNodeMockGetSignatureVerifierExpectation
	expectations       []*BaseNodeMockGetSignatureVerifierExpectation
}

// BaseNodeMockGetSignatureVerifierExpectation specifies expectation struct of the BaseNode.GetSignatureVerifier
type BaseNodeMockGetSignatureVerifierExpectation struct {
	mock *BaseNodeMock

	results *BaseNodeMockGetSignatureVerifierResults
	Counter uint64
}

// BaseNodeMockGetSignatureVerifierResults contains results of the BaseNode.GetSignatureVerifier
type BaseNodeMockGetSignatureVerifierResults struct {
	s1 cryptkit.SignatureVerifier
}

// Expect sets up expected params for BaseNode.GetSignatureVerifier
func (mmGetSignatureVerifier *mBaseNodeMockGetSignatureVerifier) Expect() *mBaseNodeMockGetSignatureVerifier {
	if mmGetSignatureVerifier.mock.funcGetSignatureVerifier != nil {
		mmGetSignatureVerifier.mock.t.Fatalf("BaseNodeMock.GetSignatureVerifier mock is already set by Set")
	}

	if mmGetSignatureVerifier.defaultExpectation == nil {
		mmGetSignatureVerifier.defaultExpectation = &BaseNodeMockGetSignatureVerifierExpectation{}
	}

	return mmGetSignatureVerifier
}

// Inspect accepts an inspector function that has same arguments as the BaseNode.GetSignatureVerifier
func (mmGetSignatureVerifier *mBaseNodeMockGetSignatureVerifier) Inspect(f func()) *mBaseNodeMockGetSignatureVerifier {
	if mmGetSignatureVerifier.mock.inspectFuncGetSignatureVerifier != nil {
		mmGetSignatureVerifier.mock.t.Fatalf("Inspect function is already set for BaseNodeMock.GetSignatureVerifier")
	}

	mmGetSignatureVerifier.mock.inspectFuncGetSignatureVerifier = f

	return mmGetSignatureVerifier
}

// Return sets up results that will be returned by BaseNode.GetSignatureVerifier
func (mmGetSignatureVerifier *mBaseNodeMockGetSignatureVerifier) Return(s1 cryptkit.SignatureVerifier) *BaseNodeMock {
	if mmGetSignatureVerifier.mock.funcGetSignatureVerifier != nil {
		mmGetSignatureVerifier.mock.t.Fatalf("BaseNodeMock.GetSignatureVerifier mock is already set by Set")
	}

	if mmGetSignatureVerifier.defaultExpectation == nil {
		mmGetSignatureVerifier.defaultExpectation = &BaseNodeMockGetSignatureVerifierExpectation{mock: mmGetSignatureVerifier.mock}
	}
	mmGetSignatureVerifier.defaultExpectation.results = &BaseNodeMockGetSignatureVerifierResults{s1}
	return mmGetSignatureVerifier.mock
}

//Set uses given function f to mock the BaseNode.GetSignatureVerifier method
func (mmGetSignatureVerifier *mBaseNodeMockGetSignatureVerifier) Set(f func() (s1 cryptkit.SignatureVerifier)) *BaseNodeMock {
	if mmGetSignatureVerifier.defaultExpectation != nil {
		mmGetSignatureVerifier.mock.t.Fatalf("Default expectation is already set for the BaseNode.GetSignatureVerifier method")
	}

	if len(mmGetSignatureVerifier.expectations) > 0 {
		mmGetSignatureVerifier.mock.t.Fatalf("Some expectations are already set for the BaseNode.GetSignatureVerifier method")
	}

	mmGetSignatureVerifier.mock.funcGetSignatureVerifier = f
	return mmGetSignatureVerifier.mock
}

// GetSignatureVerifier implements BaseNode
func (mmGetSignatureVerifier *BaseNodeMock) GetSignatureVerifier() (s1 cryptkit.SignatureVerifier) {
	mm_atomic.AddUint64(&mmGetSignatureVerifier.beforeGetSignatureVerifierCounter, 1)
	defer mm_atomic.AddUint64(&mmGetSignatureVerifier.afterGetSignatureVerifierCounter, 1)

	if mmGetSignatureVerifier.inspectFuncGetSignatureVerifier != nil {
		mmGetSignatureVerifier.inspectFuncGetSignatureVerifier()
	}

	if mmGetSignatureVerifier.GetSignatureVerifierMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetSignatureVerifier.GetSignatureVerifierMock.defaultExpectation.Counter, 1)

		mm_results := mmGetSignatureVerifier.GetSignatureVerifierMock.defaultExpectation.results
		if mm_results == nil {
			mmGetSignatureVerifier.t.Fatal("No results are set for the BaseNodeMock.GetSignatureVerifier")
		}
		return (*mm_results).s1
	}
	if mmGetSignatureVerifier.funcGetSignatureVerifier != nil {
		return mmGetSignatureVerifier.funcGetSignatureVerifier()
	}
	mmGetSignatureVerifier.t.Fatalf("Unexpected call to BaseNodeMock.GetSignatureVerifier.")
	return
}

// GetSignatureVerifierAfterCounter returns a count of finished BaseNodeMock.GetSignatureVerifier invocations
func (mmGetSignatureVerifier *BaseNodeMock) GetSignatureVerifierAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetSignatureVerifier.afterGetSignatureVerifierCounter)
}

// GetSignatureVerifierBeforeCounter returns a count of BaseNodeMock.GetSignatureVerifier invocations
func (mmGetSignatureVerifier *BaseNodeMock) GetSignatureVerifierBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetSignatureVerifier.beforeGetSignatureVerifierCounter)
}

// MinimockGetSignatureVerifierDone returns true if the count of the GetSignatureVerifier invocations corresponds
// the number of defined expectations
func (m *BaseNodeMock) MinimockGetSignatureVerifierDone() bool {
	for _, e := range m.GetSignatureVerifierMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetSignatureVerifierMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetSignatureVerifierCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetSignatureVerifier != nil && mm_atomic.LoadUint64(&m.afterGetSignatureVerifierCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetSignatureVerifierInspect logs each unmet expectation
func (m *BaseNodeMock) MinimockGetSignatureVerifierInspect() {
	for _, e := range m.GetSignatureVerifierMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to BaseNodeMock.GetSignatureVerifier")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetSignatureVerifierMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetSignatureVerifierCounter) < 1 {
		m.t.Error("Expected call to BaseNodeMock.GetSignatureVerifier")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetSignatureVerifier != nil && mm_atomic.LoadUint64(&m.afterGetSignatureVerifierCounter) < 1 {
		m.t.Error("Expected call to BaseNodeMock.GetSignatureVerifier")
	}
}

type mBaseNodeMockGetStatic struct {
	mock               *BaseNodeMock
	defaultExpectation *BaseNodeMockGetStaticExpectation
	expectations       []*BaseNodeMockGetStaticExpectation
}

// BaseNodeMockGetStaticExpectation specifies expectation struct of the BaseNode.GetStatic
type BaseNodeMockGetStaticExpectation struct {
	mock *BaseNodeMock

	results *BaseNodeMockGetStaticResults
	Counter uint64
}

// BaseNodeMockGetStaticResults contains results of the BaseNode.GetStatic
type BaseNodeMockGetStaticResults struct {
	s1 StaticProfile
}

// Expect sets up expected params for BaseNode.GetStatic
func (mmGetStatic *mBaseNodeMockGetStatic) Expect() *mBaseNodeMockGetStatic {
	if mmGetStatic.mock.funcGetStatic != nil {
		mmGetStatic.mock.t.Fatalf("BaseNodeMock.GetStatic mock is already set by Set")
	}

	if mmGetStatic.defaultExpectation == nil {
		mmGetStatic.defaultExpectation = &BaseNodeMockGetStaticExpectation{}
	}

	return mmGetStatic
}

// Inspect accepts an inspector function that has same arguments as the BaseNode.GetStatic
func (mmGetStatic *mBaseNodeMockGetStatic) Inspect(f func()) *mBaseNodeMockGetStatic {
	if mmGetStatic.mock.inspectFuncGetStatic != nil {
		mmGetStatic.mock.t.Fatalf("Inspect function is already set for BaseNodeMock.GetStatic")
	}

	mmGetStatic.mock.inspectFuncGetStatic = f

	return mmGetStatic
}

// Return sets up results that will be returned by BaseNode.GetStatic
func (mmGetStatic *mBaseNodeMockGetStatic) Return(s1 StaticProfile) *BaseNodeMock {
	if mmGetStatic.mock.funcGetStatic != nil {
		mmGetStatic.mock.t.Fatalf("BaseNodeMock.GetStatic mock is already set by Set")
	}

	if mmGetStatic.defaultExpectation == nil {
		mmGetStatic.defaultExpectation = &BaseNodeMockGetStaticExpectation{mock: mmGetStatic.mock}
	}
	mmGetStatic.defaultExpectation.results = &BaseNodeMockGetStaticResults{s1}
	return mmGetStatic.mock
}

//Set uses given function f to mock the BaseNode.GetStatic method
func (mmGetStatic *mBaseNodeMockGetStatic) Set(f func() (s1 StaticProfile)) *BaseNodeMock {
	if mmGetStatic.defaultExpectation != nil {
		mmGetStatic.mock.t.Fatalf("Default expectation is already set for the BaseNode.GetStatic method")
	}

	if len(mmGetStatic.expectations) > 0 {
		mmGetStatic.mock.t.Fatalf("Some expectations are already set for the BaseNode.GetStatic method")
	}

	mmGetStatic.mock.funcGetStatic = f
	return mmGetStatic.mock
}

// GetStatic implements BaseNode
func (mmGetStatic *BaseNodeMock) GetStatic() (s1 StaticProfile) {
	mm_atomic.AddUint64(&mmGetStatic.beforeGetStaticCounter, 1)
	defer mm_atomic.AddUint64(&mmGetStatic.afterGetStaticCounter, 1)

	if mmGetStatic.inspectFuncGetStatic != nil {
		mmGetStatic.inspectFuncGetStatic()
	}

	if mmGetStatic.GetStaticMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetStatic.GetStaticMock.defaultExpectation.Counter, 1)

		mm_results := mmGetStatic.GetStaticMock.defaultExpectation.results
		if mm_results == nil {
			mmGetStatic.t.Fatal("No results are set for the BaseNodeMock.GetStatic")
		}
		return (*mm_results).s1
	}
	if mmGetStatic.funcGetStatic != nil {
		return mmGetStatic.funcGetStatic()
	}
	mmGetStatic.t.Fatalf("Unexpected call to BaseNodeMock.GetStatic.")
	return
}

// GetStaticAfterCounter returns a count of finished BaseNodeMock.GetStatic invocations
func (mmGetStatic *BaseNodeMock) GetStaticAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetStatic.afterGetStaticCounter)
}

// GetStaticBeforeCounter returns a count of BaseNodeMock.GetStatic invocations
func (mmGetStatic *BaseNodeMock) GetStaticBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetStatic.beforeGetStaticCounter)
}

// MinimockGetStaticDone returns true if the count of the GetStatic invocations corresponds
// the number of defined expectations
func (m *BaseNodeMock) MinimockGetStaticDone() bool {
	for _, e := range m.GetStaticMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetStaticMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetStaticCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetStatic != nil && mm_atomic.LoadUint64(&m.afterGetStaticCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetStaticInspect logs each unmet expectation
func (m *BaseNodeMock) MinimockGetStaticInspect() {
	for _, e := range m.GetStaticMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to BaseNodeMock.GetStatic")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetStaticMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetStaticCounter) < 1 {
		m.t.Error("Expected call to BaseNodeMock.GetStatic")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetStatic != nil && mm_atomic.LoadUint64(&m.afterGetStaticCounter) < 1 {
		m.t.Error("Expected call to BaseNodeMock.GetStatic")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *BaseNodeMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetNodeIDInspect()

		m.MinimockGetOpModeInspect()

		m.MinimockGetSignatureVerifierInspect()

		m.MinimockGetStaticInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *BaseNodeMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *BaseNodeMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetNodeIDDone() &&
		m.MinimockGetOpModeDone() &&
		m.MinimockGetSignatureVerifierDone() &&
		m.MinimockGetStaticDone()
}
