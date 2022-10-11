package rolebinding

import (
	"captain/pkg/bussiness/kube-resources/alpha1"
	"captain/pkg/unify/query"
	"captain/pkg/unify/response"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
)

type rolebindingProvider struct {
	informers informers.SharedInformerFactory
}

func New(informer informers.SharedInformerFactory) rolebindingProvider {
	return rolebindingProvider{informers: informer}
}

func (cr rolebindingProvider) Get(namespace, name string) (runtime.Object, error) {
	return cr.informers.Rbac().V1().RoleBindings().Lister().RoleBindings(namespace).Get(name)

}

func (cr rolebindingProvider) List(namespace string, query *query.QueryInfo) (*response.ListResult, error) {
	var roleBindings []*rbacv1.RoleBinding
	var err error

	roleBindings, err = cr.informers.Rbac().V1().RoleBindings().Lister().RoleBindings(namespace).List(query.GetSelector())

	if err != nil {
		return nil, err
	}

	var result []runtime.Object
	for _, roleBinding := range roleBindings {
		result = append(result, roleBinding)
	}

	return alpha1.DefaultList(result, query, compareFunc, filter), nil
}

func filter(object runtime.Object, filter query.Filter) bool {
	role, ok := object.(*rbacv1.RoleBinding)

	if !ok {
		return false
	}
	return alpha1.DefaultObjectMetaFilter(role.ObjectMeta, filter)
}

func compareFunc(left, right runtime.Object, field query.Field) bool {

	leftRoleBinding, ok := left.(*rbacv1.RoleBinding)
	if !ok {
		return false
	}

	rightRoleBinding, ok := right.(*rbacv1.RoleBinding)
	if !ok {
		return false
	}

	return alpha1.DefaultObjectMetaCompare(leftRoleBinding.ObjectMeta, rightRoleBinding.ObjectMeta, field)
}
