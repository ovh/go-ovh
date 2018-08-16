package metrics

// Add a new label
func (labels TokenLabels) Add(label, value string) TokenLabels {
	if labels == nil || label == "" || value == "" {
		return labels
	}

	return append(labels, TokenLabel{
		Key:   label,
		Value: value,
	})
}

// Remove existing label
func (labels TokenLabels) Remove(label string) TokenLabels {
	if labels == nil || len(labels) == 0 {
		return labels
	}

	for i, l := range labels {
		if l.Key == label {
			return append(labels[:i], labels[i+1:]...)
		}
	}
	return labels
}

// Len implement sort.Sort()
func (labels TokenLabels) Len() int {
	return len(labels)
}

// Swap implement sort.Sort()
func (labels TokenLabels) Swap(i, j int) {
	labels[i], labels[j] = labels[j], labels[i]
}

// Less implement sort.Sort()
func (labels TokenLabels) Less(i, j int) bool {
	return labels[i].Key < labels[j].Key
}
