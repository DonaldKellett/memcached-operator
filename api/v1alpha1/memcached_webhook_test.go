/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Memcached Webhook", func() {

	Context("When creating Memcached under Defaulting Webhook", func() {
		It("Should fill in the default value if a required field is empty", func() {
			var m *Memcached = &Memcached{}
			var defaultSize, defaultContainerPort int32 = 1, 11211
			m.Default()
			Expect(m.Spec.Size).To(Equal(defaultSize))
			Expect(m.Spec.ContainerPort).To(Equal(defaultContainerPort))
		})
	})

	Context("When creating Memcached under Validating Webhook", func() {
		It("Should allow creation of valid Memcached CR", func() {
			var m *Memcached = &Memcached{Spec: MemcachedSpec{Size: 3, ContainerPort: 11211}}
			warn, err := m.ValidateCreate()
			Expect(warn).To(BeNil())
			Expect(err).To(BeNil())
		})

		It("Should disallow a negative size", func() {
			var m *Memcached = &Memcached{Spec: MemcachedSpec{Size: -1, ContainerPort: 11211}}
			_, err := m.ValidateCreate()
			Expect(err).To(MatchError(ContainSubstring("must be between 1 and 5, both inclusive")))
		})

		It("Should disallow a size greater than 5", func() {
			var m *Memcached = &Memcached{Spec: MemcachedSpec{Size: 6, ContainerPort: 11211}}
			_, err := m.ValidateCreate()
			Expect(err).To(MatchError(ContainSubstring("must be between 1 and 5, both inclusive")))
		})

		It("Should disallow a negative containerPort", func() {
			var m *Memcached = &Memcached{Spec: MemcachedSpec{Size: 3, ContainerPort: -1}}
			_, err := m.ValidateCreate()
			Expect(err).To(MatchError(ContainSubstring("must be between 0 and 65536, both exclusive")))
		})

		It("Should disallow a containerPort greater than 65535", func() {
			var m *Memcached = &Memcached{Spec: MemcachedSpec{Size: 3, ContainerPort: 65536}}
			_, err := m.ValidateCreate()
			Expect(err).To(MatchError(ContainSubstring("must be between 0 and 65536, both exclusive")))
		})

		It("Should allow update of a valid Memcached CR", func() {
			var m *Memcached = &Memcached{Spec: MemcachedSpec{Size: 5, ContainerPort: 11211}}
			var old *Memcached = &Memcached{Spec: MemcachedSpec{Size: 3, ContainerPort: 11211}}
			warn, err := m.ValidateUpdate(old)
			Expect(warn).To(BeNil())
			Expect(err).To(BeNil())
		})

		It("Should disallow an update to containerPort", func() {
			var m *Memcached = &Memcached{Spec: MemcachedSpec{Size: 3, ContainerPort: 33221}}
			var old *Memcached = &Memcached{Spec: MemcachedSpec{Size: 3, ContainerPort: 11211}}
			_, err := m.ValidateUpdate(old)
			Expect(err).To(MatchError(ContainSubstring("containerPort field cannot be modified after creation")))
		})
	})

})
