// This file was generated by counterfeiter
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeOrgActor struct {
	GetOrganizationByNameStub        func(orgName string) (v2action.Organization, v2action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		orgName string
	}
	getOrganizationByNameReturns struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	GetOrganizationSummaryByNameStub        func(orgName string) (v2action.OrganizationSummary, v2action.Warnings, error)
	getOrganizationSummaryByNameMutex       sync.RWMutex
	getOrganizationSummaryByNameArgsForCall []struct {
		orgName string
	}
	getOrganizationSummaryByNameReturns struct {
		result1 v2action.OrganizationSummary
		result2 v2action.Warnings
		result3 error
	}
	getOrganizationSummaryByNameReturnsOnCall map[int]struct {
		result1 v2action.OrganizationSummary
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrgActor) GetOrganizationByName(orgName string) (v2action.Organization, v2action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("GetOrganizationByName", []interface{}{orgName})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(orgName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrganizationByNameReturns.result1, fake.getOrganizationByNameReturns.result2, fake.getOrganizationByNameReturns.result3
}

func (fake *FakeOrgActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeOrgActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return fake.getOrganizationByNameArgsForCall[i].orgName
}

func (fake *FakeOrgActor) GetOrganizationByNameReturns(result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgActor) GetOrganizationByNameReturnsOnCall(i int, result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Organization
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgActor) GetOrganizationSummaryByName(orgName string) (v2action.OrganizationSummary, v2action.Warnings, error) {
	fake.getOrganizationSummaryByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationSummaryByNameReturnsOnCall[len(fake.getOrganizationSummaryByNameArgsForCall)]
	fake.getOrganizationSummaryByNameArgsForCall = append(fake.getOrganizationSummaryByNameArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("GetOrganizationSummaryByName", []interface{}{orgName})
	fake.getOrganizationSummaryByNameMutex.Unlock()
	if fake.GetOrganizationSummaryByNameStub != nil {
		return fake.GetOrganizationSummaryByNameStub(orgName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrganizationSummaryByNameReturns.result1, fake.getOrganizationSummaryByNameReturns.result2, fake.getOrganizationSummaryByNameReturns.result3
}

func (fake *FakeOrgActor) GetOrganizationSummaryByNameCallCount() int {
	fake.getOrganizationSummaryByNameMutex.RLock()
	defer fake.getOrganizationSummaryByNameMutex.RUnlock()
	return len(fake.getOrganizationSummaryByNameArgsForCall)
}

func (fake *FakeOrgActor) GetOrganizationSummaryByNameArgsForCall(i int) string {
	fake.getOrganizationSummaryByNameMutex.RLock()
	defer fake.getOrganizationSummaryByNameMutex.RUnlock()
	return fake.getOrganizationSummaryByNameArgsForCall[i].orgName
}

func (fake *FakeOrgActor) GetOrganizationSummaryByNameReturns(result1 v2action.OrganizationSummary, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationSummaryByNameStub = nil
	fake.getOrganizationSummaryByNameReturns = struct {
		result1 v2action.OrganizationSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgActor) GetOrganizationSummaryByNameReturnsOnCall(i int, result1 v2action.OrganizationSummary, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationSummaryByNameStub = nil
	if fake.getOrganizationSummaryByNameReturnsOnCall == nil {
		fake.getOrganizationSummaryByNameReturnsOnCall = make(map[int]struct {
			result1 v2action.OrganizationSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrganizationSummaryByNameReturnsOnCall[i] = struct {
		result1 v2action.OrganizationSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getOrganizationSummaryByNameMutex.RLock()
	defer fake.getOrganizationSummaryByNameMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeOrgActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v2.OrgActor = new(FakeOrgActor)
