package openshift

import lokiv1 "github.com/agardiman/loki/operator/apis/loki/v1"

func RecordingRuleTenantLabels(r *lokiv1.RecordingRule) {
	switch r.Spec.TenantID {
	case tenantApplication:
		for groupIdx, group := range r.Spec.Groups {
			group := group
			for ruleIdx, rule := range group.Rules {
				rule := rule
				if rule.Labels == nil {
					rule.Labels = map[string]string{}
				}
				rule.Labels[opaDefaultLabelMatcher] = r.Namespace
				group.Rules[ruleIdx] = rule
			}
			r.Spec.Groups[groupIdx] = group
		}
	case tenantInfrastructure, tenantAudit:
		// Do nothing
	case tenantNetwork:
		// Do nothing
	default:
		// Do nothing
	}
}
