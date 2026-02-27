package utils

import (
	"testing"

	apps "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	workloadsv1alpha1 "sigs.k8s.io/rbgs/api/workloads/v1alpha1"
)

func TestGetRoleReplicaIndex(t *testing.T) {
	tests := []struct {
		name     string
		instance *workloadsv1alpha1.Instance
		expect   string
	}{
		{
			name: "use pod-index label when present",
			instance: &workloadsv1alpha1.Instance{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo-set-3",
					Labels: map[string]string{
						apps.PodIndexLabel:           "3",
						apps.StatefulSetPodNameLabel: "demo-set-3",
					},
				},
			},
			expect: "3",
		},
		{
			name: "fallback to parse stateful instance name",
			instance: &workloadsv1alpha1.Instance{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo-set-7",
					Labels: map[string]string{
						apps.StatefulSetPodNameLabel: "demo-set-7",
					},
				},
			},
			expect: "7",
		},
		{
			name: "return empty for stateless instance",
			instance: &workloadsv1alpha1.Instance{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo-set-ab12c",
					Labels: map[string]string{
						workloadsv1alpha1.SetInstanceIDLabelKey: "ab12c",
					},
				},
			},
			expect: "",
		},
		{
			name: "return empty for invalid ordinal suffix",
			instance: &workloadsv1alpha1.Instance{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo-set-abc",
					Labels: map[string]string{
						apps.StatefulSetPodNameLabel: "demo-set-abc",
					},
				},
			},
			expect: "",
		},
		{
			name:     "return empty for nil instance",
			instance: nil,
			expect:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRoleReplicaIndex(tt.instance)
			if got != tt.expect {
				t.Fatalf("GetRoleReplicaIndex() = %q, want %q", got, tt.expect)
			}
		})
	}
}
