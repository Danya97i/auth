// Code generated by http://github.com/gojuno/minimock (v3.4.2). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Danya97i/auth/internal/repository.UserCache -o user_cache_minimock.go -n UserCacheMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/Danya97i/auth/internal/models"
	"github.com/gojuno/minimock/v3"
)

// UserCacheMock implements repository.UserCache
type UserCacheMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGet          func(ctx context.Context, id int64) (up1 *models.User, err error)
	inspectFuncGet   func(ctx context.Context, id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mUserCacheMockGet

	funcSet          func(ctx context.Context, user *models.User) (err error)
	inspectFuncSet   func(ctx context.Context, user *models.User)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mUserCacheMockSet
}

// NewUserCacheMock returns a mock for repository.UserCache
func NewUserCacheMock(t minimock.Tester) *UserCacheMock {
	m := &UserCacheMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetMock = mUserCacheMockGet{mock: m}
	m.GetMock.callArgs = []*UserCacheMockGetParams{}

	m.SetMock = mUserCacheMockSet{mock: m}
	m.SetMock.callArgs = []*UserCacheMockSetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserCacheMockGet struct {
	optional           bool
	mock               *UserCacheMock
	defaultExpectation *UserCacheMockGetExpectation
	expectations       []*UserCacheMockGetExpectation

	callArgs []*UserCacheMockGetParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserCacheMockGetExpectation specifies expectation struct of the UserCache.Get
type UserCacheMockGetExpectation struct {
	mock      *UserCacheMock
	params    *UserCacheMockGetParams
	paramPtrs *UserCacheMockGetParamPtrs
	results   *UserCacheMockGetResults
	Counter   uint64
}

// UserCacheMockGetParams contains parameters of the UserCache.Get
type UserCacheMockGetParams struct {
	ctx context.Context
	id  int64
}

// UserCacheMockGetParamPtrs contains pointers to parameters of the UserCache.Get
type UserCacheMockGetParamPtrs struct {
	ctx *context.Context
	id  *int64
}

// UserCacheMockGetResults contains results of the UserCache.Get
type UserCacheMockGetResults struct {
	up1 *models.User
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGet *mUserCacheMockGet) Optional() *mUserCacheMockGet {
	mmGet.optional = true
	return mmGet
}

// Expect sets up expected params for UserCache.Get
func (mmGet *mUserCacheMockGet) Expect(ctx context.Context, id int64) *mUserCacheMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserCacheMockGetExpectation{}
	}

	if mmGet.defaultExpectation.paramPtrs != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by ExpectParams functions")
	}

	mmGet.defaultExpectation.params = &UserCacheMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// ExpectCtxParam1 sets up expected param ctx for UserCache.Get
func (mmGet *mUserCacheMockGet) ExpectCtxParam1(ctx context.Context) *mUserCacheMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserCacheMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &UserCacheMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.ctx = &ctx

	return mmGet
}

// ExpectIdParam2 sets up expected param id for UserCache.Get
func (mmGet *mUserCacheMockGet) ExpectIdParam2(id int64) *mUserCacheMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserCacheMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &UserCacheMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.id = &id

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the UserCache.Get
func (mmGet *mUserCacheMockGet) Inspect(f func(ctx context.Context, id int64)) *mUserCacheMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for UserCacheMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by UserCache.Get
func (mmGet *mUserCacheMockGet) Return(up1 *models.User, err error) *UserCacheMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserCacheMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &UserCacheMockGetResults{up1, err}
	return mmGet.mock
}

// Set uses given function f to mock the UserCache.Get method
func (mmGet *mUserCacheMockGet) Set(f func(ctx context.Context, id int64) (up1 *models.User, err error)) *UserCacheMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the UserCache.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the UserCache.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the UserCache.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mUserCacheMockGet) When(ctx context.Context, id int64) *UserCacheMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	expectation := &UserCacheMockGetExpectation{
		mock:   mmGet.mock,
		params: &UserCacheMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up UserCache.Get return parameters for the expectation previously defined by the When method
func (e *UserCacheMockGetExpectation) Then(up1 *models.User, err error) *UserCacheMock {
	e.results = &UserCacheMockGetResults{up1, err}
	return e.mock
}

// Times sets number of times UserCache.Get should be invoked
func (mmGet *mUserCacheMockGet) Times(n uint64) *mUserCacheMockGet {
	if n == 0 {
		mmGet.mock.t.Fatalf("Times of UserCacheMock.Get mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGet.expectedInvocations, n)
	return mmGet
}

func (mmGet *mUserCacheMockGet) invocationsDone() bool {
	if len(mmGet.expectations) == 0 && mmGet.defaultExpectation == nil && mmGet.mock.funcGet == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGet.mock.afterGetCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGet.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Get implements repository.UserCache
func (mmGet *UserCacheMock) Get(ctx context.Context, id int64) (up1 *models.User, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := UserCacheMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_want_ptrs := mmGet.GetMock.defaultExpectation.paramPtrs

		mm_got := UserCacheMockGetParams{ctx, id}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGet.t.Errorf("UserCacheMock.Get got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmGet.t.Errorf("UserCacheMock.Get got unexpected parameter id, want: %#v, got: %#v%s\n", *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("UserCacheMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the UserCacheMock.Get")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to UserCacheMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished UserCacheMock.Get invocations
func (mmGet *UserCacheMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of UserCacheMock.Get invocations
func (mmGet *UserCacheMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to UserCacheMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mUserCacheMockGet) Calls() []*UserCacheMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*UserCacheMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *UserCacheMock) MinimockGetDone() bool {
	if m.GetMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetMock.invocationsDone()
}

// MinimockGetInspect logs each unmet expectation
func (m *UserCacheMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserCacheMock.Get with params: %#v", *e.params)
		}
	}

	afterGetCounter := mm_atomic.LoadUint64(&m.afterGetCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && afterGetCounter < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserCacheMock.Get")
		} else {
			m.t.Errorf("Expected call to UserCacheMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && afterGetCounter < 1 {
		m.t.Error("Expected call to UserCacheMock.Get")
	}

	if !m.GetMock.invocationsDone() && afterGetCounter > 0 {
		m.t.Errorf("Expected %d calls to UserCacheMock.Get but found %d calls",
			mm_atomic.LoadUint64(&m.GetMock.expectedInvocations), afterGetCounter)
	}
}

type mUserCacheMockSet struct {
	optional           bool
	mock               *UserCacheMock
	defaultExpectation *UserCacheMockSetExpectation
	expectations       []*UserCacheMockSetExpectation

	callArgs []*UserCacheMockSetParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserCacheMockSetExpectation specifies expectation struct of the UserCache.Set
type UserCacheMockSetExpectation struct {
	mock      *UserCacheMock
	params    *UserCacheMockSetParams
	paramPtrs *UserCacheMockSetParamPtrs
	results   *UserCacheMockSetResults
	Counter   uint64
}

// UserCacheMockSetParams contains parameters of the UserCache.Set
type UserCacheMockSetParams struct {
	ctx  context.Context
	user *models.User
}

// UserCacheMockSetParamPtrs contains pointers to parameters of the UserCache.Set
type UserCacheMockSetParamPtrs struct {
	ctx  *context.Context
	user **models.User
}

// UserCacheMockSetResults contains results of the UserCache.Set
type UserCacheMockSetResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSet *mUserCacheMockSet) Optional() *mUserCacheMockSet {
	mmSet.optional = true
	return mmSet
}

// Expect sets up expected params for UserCache.Set
func (mmSet *mUserCacheMockSet) Expect(ctx context.Context, user *models.User) *mUserCacheMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("UserCacheMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &UserCacheMockSetExpectation{}
	}

	if mmSet.defaultExpectation.paramPtrs != nil {
		mmSet.mock.t.Fatalf("UserCacheMock.Set mock is already set by ExpectParams functions")
	}

	mmSet.defaultExpectation.params = &UserCacheMockSetParams{ctx, user}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// ExpectCtxParam1 sets up expected param ctx for UserCache.Set
func (mmSet *mUserCacheMockSet) ExpectCtxParam1(ctx context.Context) *mUserCacheMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("UserCacheMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &UserCacheMockSetExpectation{}
	}

	if mmSet.defaultExpectation.params != nil {
		mmSet.mock.t.Fatalf("UserCacheMock.Set mock is already set by Expect")
	}

	if mmSet.defaultExpectation.paramPtrs == nil {
		mmSet.defaultExpectation.paramPtrs = &UserCacheMockSetParamPtrs{}
	}
	mmSet.defaultExpectation.paramPtrs.ctx = &ctx

	return mmSet
}

// ExpectUserParam2 sets up expected param user for UserCache.Set
func (mmSet *mUserCacheMockSet) ExpectUserParam2(user *models.User) *mUserCacheMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("UserCacheMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &UserCacheMockSetExpectation{}
	}

	if mmSet.defaultExpectation.params != nil {
		mmSet.mock.t.Fatalf("UserCacheMock.Set mock is already set by Expect")
	}

	if mmSet.defaultExpectation.paramPtrs == nil {
		mmSet.defaultExpectation.paramPtrs = &UserCacheMockSetParamPtrs{}
	}
	mmSet.defaultExpectation.paramPtrs.user = &user

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the UserCache.Set
func (mmSet *mUserCacheMockSet) Inspect(f func(ctx context.Context, user *models.User)) *mUserCacheMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for UserCacheMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by UserCache.Set
func (mmSet *mUserCacheMockSet) Return(err error) *UserCacheMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("UserCacheMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &UserCacheMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &UserCacheMockSetResults{err}
	return mmSet.mock
}

// Set uses given function f to mock the UserCache.Set method
func (mmSet *mUserCacheMockSet) Set(f func(ctx context.Context, user *models.User) (err error)) *UserCacheMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the UserCache.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the UserCache.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the UserCache.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mUserCacheMockSet) When(ctx context.Context, user *models.User) *UserCacheMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("UserCacheMock.Set mock is already set by Set")
	}

	expectation := &UserCacheMockSetExpectation{
		mock:   mmSet.mock,
		params: &UserCacheMockSetParams{ctx, user},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up UserCache.Set return parameters for the expectation previously defined by the When method
func (e *UserCacheMockSetExpectation) Then(err error) *UserCacheMock {
	e.results = &UserCacheMockSetResults{err}
	return e.mock
}

// Times sets number of times UserCache.Set should be invoked
func (mmSet *mUserCacheMockSet) Times(n uint64) *mUserCacheMockSet {
	if n == 0 {
		mmSet.mock.t.Fatalf("Times of UserCacheMock.Set mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSet.expectedInvocations, n)
	return mmSet
}

func (mmSet *mUserCacheMockSet) invocationsDone() bool {
	if len(mmSet.expectations) == 0 && mmSet.defaultExpectation == nil && mmSet.mock.funcSet == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSet.mock.afterSetCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSet.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Set implements repository.UserCache
func (mmSet *UserCacheMock) Set(ctx context.Context, user *models.User) (err error) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(ctx, user)
	}

	mm_params := UserCacheMockSetParams{ctx, user}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, &mm_params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		mm_want := mmSet.SetMock.defaultExpectation.params
		mm_want_ptrs := mmSet.SetMock.defaultExpectation.paramPtrs

		mm_got := UserCacheMockSetParams{ctx, user}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmSet.t.Errorf("UserCacheMock.Set got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.user != nil && !minimock.Equal(*mm_want_ptrs.user, mm_got.user) {
				mmSet.t.Errorf("UserCacheMock.Set got unexpected parameter user, want: %#v, got: %#v%s\n", *mm_want_ptrs.user, mm_got.user, minimock.Diff(*mm_want_ptrs.user, mm_got.user))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSet.t.Errorf("UserCacheMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSet.SetMock.defaultExpectation.results
		if mm_results == nil {
			mmSet.t.Fatal("No results are set for the UserCacheMock.Set")
		}
		return (*mm_results).err
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(ctx, user)
	}
	mmSet.t.Fatalf("Unexpected call to UserCacheMock.Set. %v %v", ctx, user)
	return
}

// SetAfterCounter returns a count of finished UserCacheMock.Set invocations
func (mmSet *UserCacheMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of UserCacheMock.Set invocations
func (mmSet *UserCacheMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to UserCacheMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mUserCacheMockSet) Calls() []*UserCacheMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*UserCacheMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *UserCacheMock) MinimockSetDone() bool {
	if m.SetMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SetMock.invocationsDone()
}

// MinimockSetInspect logs each unmet expectation
func (m *UserCacheMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserCacheMock.Set with params: %#v", *e.params)
		}
	}

	afterSetCounter := mm_atomic.LoadUint64(&m.afterSetCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && afterSetCounter < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserCacheMock.Set")
		} else {
			m.t.Errorf("Expected call to UserCacheMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && afterSetCounter < 1 {
		m.t.Error("Expected call to UserCacheMock.Set")
	}

	if !m.SetMock.invocationsDone() && afterSetCounter > 0 {
		m.t.Errorf("Expected %d calls to UserCacheMock.Set but found %d calls",
			mm_atomic.LoadUint64(&m.SetMock.expectedInvocations), afterSetCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserCacheMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetInspect()

			m.MinimockSetInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserCacheMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *UserCacheMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetDone() &&
		m.MinimockSetDone()
}