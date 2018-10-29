package v1alpha1

// Automatically generated file, do not edit!

/*
This file contains methods on the types in the v1alpha1 package for
determining resource names.

Resource names are specified with the '+freeze-api:resource-name' special
comment, on types containing meta.TypeMeta. Resource names are specified
statically, and do not change at runtime.
*/

// ResourceName returns the resource name for a ClusterRole.
// The resource name for ClusterRole is "clusterRoles".
func (r ClusterRole) ResourceName() string {
	return "clusterRoles"
}

// ResourceName returns the resource name for a ClusterRoleBinding.
// The resource name for ClusterRoleBinding is "clusterRoleBindings".
func (r ClusterRoleBinding) ResourceName() string {
	return "clusterRoleBindings"
}

// ResourceName returns the resource name for a Role.
// The resource name for Role is "roles".
func (r Role) ResourceName() string {
	return "roles"
}

// ResourceName returns the resource name for a RoleBinding.
// The resource name for RoleBinding is "roleBindings".
func (r RoleBinding) ResourceName() string {
	return "roleBindings"
}

// ResourceName returns the resource name for a Subject.
// The resource name for Subject is "subjects".
func (r Subject) ResourceName() string {
	return "subjects"
}