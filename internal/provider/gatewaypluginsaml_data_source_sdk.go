// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginSamlDataSourceModel) RefreshFromSharedSamlPlugin(resp *shared.SamlPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateSamlPluginConfig{}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			r.Config.AssertionConsumerPath = types.StringPointerValue(resp.Config.AssertionConsumerPath)
			r.Config.IdpCertificate = types.StringPointerValue(resp.Config.IdpCertificate)
			r.Config.IdpSsoURL = types.StringPointerValue(resp.Config.IdpSsoURL)
			r.Config.Issuer = types.StringPointerValue(resp.Config.Issuer)
			if resp.Config.NameidFormat != nil {
				r.Config.NameidFormat = types.StringValue(string(*resp.Config.NameidFormat))
			} else {
				r.Config.NameidFormat = types.StringNull()
			}
			if resp.Config.RequestDigestAlgorithm != nil {
				r.Config.RequestDigestAlgorithm = types.StringValue(string(*resp.Config.RequestDigestAlgorithm))
			} else {
				r.Config.RequestDigestAlgorithm = types.StringNull()
			}
			if resp.Config.RequestSignatureAlgorithm != nil {
				r.Config.RequestSignatureAlgorithm = types.StringValue(string(*resp.Config.RequestSignatureAlgorithm))
			} else {
				r.Config.RequestSignatureAlgorithm = types.StringNull()
			}
			r.Config.RequestSigningCertificate = types.StringPointerValue(resp.Config.RequestSigningCertificate)
			r.Config.RequestSigningKey = types.StringPointerValue(resp.Config.RequestSigningKey)
			if resp.Config.ResponseDigestAlgorithm != nil {
				r.Config.ResponseDigestAlgorithm = types.StringValue(string(*resp.Config.ResponseDigestAlgorithm))
			} else {
				r.Config.ResponseDigestAlgorithm = types.StringNull()
			}
			r.Config.ResponseEncryptionKey = types.StringPointerValue(resp.Config.ResponseEncryptionKey)
			if resp.Config.ResponseSignatureAlgorithm != nil {
				r.Config.ResponseSignatureAlgorithm = types.StringValue(string(*resp.Config.ResponseSignatureAlgorithm))
			} else {
				r.Config.ResponseSignatureAlgorithm = types.StringNull()
			}
			if resp.Config.SessionAbsoluteTimeout != nil {
				r.Config.SessionAbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionAbsoluteTimeout)))
			} else {
				r.Config.SessionAbsoluteTimeout = types.NumberNull()
			}
			r.Config.SessionAudience = types.StringPointerValue(resp.Config.SessionAudience)
			r.Config.SessionCookieDomain = types.StringPointerValue(resp.Config.SessionCookieDomain)
			r.Config.SessionCookieHTTPOnly = types.BoolPointerValue(resp.Config.SessionCookieHTTPOnly)
			r.Config.SessionCookieName = types.StringPointerValue(resp.Config.SessionCookieName)
			r.Config.SessionCookiePath = types.StringPointerValue(resp.Config.SessionCookiePath)
			if resp.Config.SessionCookieSameSite != nil {
				r.Config.SessionCookieSameSite = types.StringValue(string(*resp.Config.SessionCookieSameSite))
			} else {
				r.Config.SessionCookieSameSite = types.StringNull()
			}
			r.Config.SessionCookieSecure = types.BoolPointerValue(resp.Config.SessionCookieSecure)
			r.Config.SessionEnforceSameSubject = types.BoolPointerValue(resp.Config.SessionEnforceSameSubject)
			r.Config.SessionHashStorageKey = types.BoolPointerValue(resp.Config.SessionHashStorageKey)
			r.Config.SessionHashSubject = types.BoolPointerValue(resp.Config.SessionHashSubject)
			if resp.Config.SessionIdlingTimeout != nil {
				r.Config.SessionIdlingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionIdlingTimeout)))
			} else {
				r.Config.SessionIdlingTimeout = types.NumberNull()
			}
			r.Config.SessionMemcachedHost = types.StringPointerValue(resp.Config.SessionMemcachedHost)
			r.Config.SessionMemcachedPort = types.Int64PointerValue(resp.Config.SessionMemcachedPort)
			r.Config.SessionMemcachedPrefix = types.StringPointerValue(resp.Config.SessionMemcachedPrefix)
			r.Config.SessionMemcachedSocket = types.StringPointerValue(resp.Config.SessionMemcachedSocket)
			r.Config.SessionRedisClusterMaxRedirections = types.Int64PointerValue(resp.Config.SessionRedisClusterMaxRedirections)
			r.Config.SessionRedisClusterNodes = []tfTypes.SessionRedisClusterNodes{}
			if len(r.Config.SessionRedisClusterNodes) > len(resp.Config.SessionRedisClusterNodes) {
				r.Config.SessionRedisClusterNodes = r.Config.SessionRedisClusterNodes[:len(resp.Config.SessionRedisClusterNodes)]
			}
			for sessionRedisClusterNodesCount, sessionRedisClusterNodesItem := range resp.Config.SessionRedisClusterNodes {
				var sessionRedisClusterNodes1 tfTypes.SessionRedisClusterNodes
				sessionRedisClusterNodes1.IP = types.StringPointerValue(sessionRedisClusterNodesItem.IP)
				sessionRedisClusterNodes1.Port = types.Int64PointerValue(sessionRedisClusterNodesItem.Port)
				if sessionRedisClusterNodesCount+1 > len(r.Config.SessionRedisClusterNodes) {
					r.Config.SessionRedisClusterNodes = append(r.Config.SessionRedisClusterNodes, sessionRedisClusterNodes1)
				} else {
					r.Config.SessionRedisClusterNodes[sessionRedisClusterNodesCount].IP = sessionRedisClusterNodes1.IP
					r.Config.SessionRedisClusterNodes[sessionRedisClusterNodesCount].Port = sessionRedisClusterNodes1.Port
				}
			}
			r.Config.SessionRedisConnectTimeout = types.Int64PointerValue(resp.Config.SessionRedisConnectTimeout)
			r.Config.SessionRedisHost = types.StringPointerValue(resp.Config.SessionRedisHost)
			r.Config.SessionRedisPassword = types.StringPointerValue(resp.Config.SessionRedisPassword)
			r.Config.SessionRedisPort = types.Int64PointerValue(resp.Config.SessionRedisPort)
			r.Config.SessionRedisPrefix = types.StringPointerValue(resp.Config.SessionRedisPrefix)
			r.Config.SessionRedisReadTimeout = types.Int64PointerValue(resp.Config.SessionRedisReadTimeout)
			r.Config.SessionRedisSendTimeout = types.Int64PointerValue(resp.Config.SessionRedisSendTimeout)
			r.Config.SessionRedisServerName = types.StringPointerValue(resp.Config.SessionRedisServerName)
			r.Config.SessionRedisSocket = types.StringPointerValue(resp.Config.SessionRedisSocket)
			r.Config.SessionRedisSsl = types.BoolPointerValue(resp.Config.SessionRedisSsl)
			r.Config.SessionRedisSslVerify = types.BoolPointerValue(resp.Config.SessionRedisSslVerify)
			r.Config.SessionRedisUsername = types.StringPointerValue(resp.Config.SessionRedisUsername)
			r.Config.SessionRemember = types.BoolPointerValue(resp.Config.SessionRemember)
			if resp.Config.SessionRememberAbsoluteTimeout != nil {
				r.Config.SessionRememberAbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionRememberAbsoluteTimeout)))
			} else {
				r.Config.SessionRememberAbsoluteTimeout = types.NumberNull()
			}
			r.Config.SessionRememberCookieName = types.StringPointerValue(resp.Config.SessionRememberCookieName)
			if resp.Config.SessionRememberRollingTimeout != nil {
				r.Config.SessionRememberRollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionRememberRollingTimeout)))
			} else {
				r.Config.SessionRememberRollingTimeout = types.NumberNull()
			}
			r.Config.SessionRequestHeaders = []types.String{}
			for _, v := range resp.Config.SessionRequestHeaders {
				r.Config.SessionRequestHeaders = append(r.Config.SessionRequestHeaders, types.StringValue(string(v)))
			}
			r.Config.SessionResponseHeaders = []types.String{}
			for _, v := range resp.Config.SessionResponseHeaders {
				r.Config.SessionResponseHeaders = append(r.Config.SessionResponseHeaders, types.StringValue(string(v)))
			}
			if resp.Config.SessionRollingTimeout != nil {
				r.Config.SessionRollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionRollingTimeout)))
			} else {
				r.Config.SessionRollingTimeout = types.NumberNull()
			}
			r.Config.SessionSecret = types.StringPointerValue(resp.Config.SessionSecret)
			if resp.Config.SessionStorage != nil {
				r.Config.SessionStorage = types.StringValue(string(*resp.Config.SessionStorage))
			} else {
				r.Config.SessionStorage = types.StringNull()
			}
			r.Config.SessionStoreMetadata = types.BoolPointerValue(resp.Config.SessionStoreMetadata)
			r.Config.ValidateAssertionSignature = types.BoolPointerValue(resp.Config.ValidateAssertionSignature)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
