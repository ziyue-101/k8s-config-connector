{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}IdentityPlatformTenantOAuthIDPConfig{% endblock %}
{% block body %}



Note: You must enable the Google Identity Platform in the marketplace prior to using this resource.

Note: You must specify the resource's name in <code>metadata.name</code> or <code>spec.resourceID</code> starting with "oidc.".

<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Identity Platform</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/identity-platform/docs/">/identity-platform/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v2.projects.tenants.oauthIdpConfigs</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/identity-platform/docs/reference/rest/v2/projects.tenants.oauthIdpConfigs">/identity-platform/docs/reference/rest/v2/projects.tenants.oauthIdpConfigs</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>IdentityPlatformTenantOAuthIDPConfig<br>gcpidentityplatformtenantoauthidpconfig<br>gcpidentityplatformtenantoauthidpconfigs<br>identityplatformtenantoauthidpconfig</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>identitytoolkit.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>identityplatformtenantoauthidpconfigs.identityplatform.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/project-id</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
  ```yaml
  clientId: string
  clientSecret:
    value: string
    valueFrom:
      secretKeyRef:
        key: string
        name: string
  displayName: string
  enabled: boolean
  issuer: string
  resourceID: string
  responseType:
    code: boolean
    idToken: boolean
    token: boolean
  tenantRef:
    external: string
    name: string
    namespace: string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>clientId</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The client id of an OAuth client.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clientSecret</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The client secret of the OAuth client, to enable OIDC code flow.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clientSecret.value</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Value of the field. Cannot be used if 'valueFrom' is specified.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clientSecret.valueFrom</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Source for the field's value. Cannot be used if 'value' is specified.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clientSecret.valueFrom.secretKeyRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Reference to a value with the given key in the given Secret in the resource's namespace.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clientSecret.valueFrom.secretKeyRef.key</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Key that identifies the value to be extracted.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clientSecret.valueFrom.secretKeyRef.name</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the Secret to extract a value from.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>displayName</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The config's display name set by developers.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>enabled</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}True if allows the user to sign in with the provider.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>issuer</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}For OIDC Idps, the issuer identifier.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>responseType</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The multiple response type to request for in the OAuth authorization flow. This can possibly be a combination of set bits (e.g.: {id\_token, token}).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>responseType.code</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}If true, authorization code is returned from IdP's authorization endpoint.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>responseType.idToken</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}If true, ID token is returned from IdP's authorization endpoint.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>responseType.token</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}If true, access token is returned from IdP's authorization endpoint.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>tenantRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>tenantRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The tenant for the resource

Allowed value: The Google Cloud resource name of an `IdentityPlatformTenant` resource (format: `projects/{{project}}/tenants/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>tenantRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>tenantRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>


<p>{% verbatim %}* Field is required when parent field is specified{% endverbatim %}</p>


### Status
#### Schema
  ```yaml
  conditions:
  - lastTransitionTime: string
    message: string
    reason: string
    status: string
    type: string
  observedGeneration: integer
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
  ```yaml
  # Copyright 2020 Google LLC
  #
  # Licensed under the Apache License, Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at
  #
  #     http://www.apache.org/licenses/LICENSE-2.0
  #
  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  
  apiVersion: identityplatform.cnrm.cloud.google.com/v1beta1
  kind: IdentityPlatformTenantOAuthIDPConfig
  metadata:
    labels:
      foo: bar
    name: identityplatformtenantoauthidpconfig-sample
  spec:
    resourceID: "oidc.tenant-oauth-idp-config-sample" # Must start with 'oidc.'
    tenantRef:
      name: identityplatformtenantoauthidpconfig-dep
    displayName: "sample tenant oauth idp config"
    clientId: "client-id"
    issuer: "issuer"
    enabled: true
    clientSecret:
      valueFrom:
        secretKeyRef:
          key: clientSecret
          name: identityplatformtenantoauthidpconfig-dep
  ---
  apiVersion: identityplatform.cnrm.cloud.google.com/v1beta1
  kind: IdentityPlatformTenant
  metadata:
    name: identityplatformtenantoauthidpconfig-dep
  spec:
    displayName: "test-tenant"
    allowPasswordSignup: true
    enableAnonymousUser: false
    mfaConfig:
      state: "ENABLED"
    testPhoneNumbers:
      "+12345678901": "123451"
      "+16505550000": "123450"
  ---
  apiVersion: v1
  kind: Secret
  metadata:
    name: identityplatformtenantoauthidpconfig-dep
  stringData:
    clientSecret: "secret1"
  ```


{% endblock %}