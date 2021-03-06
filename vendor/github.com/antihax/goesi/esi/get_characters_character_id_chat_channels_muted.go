/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.6.2
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"time"
)

/* A list of GetCharactersCharacterIdChatChannelsMuted. */
//easyjson:json
type GetCharactersCharacterIdChatChannelsMutedList []GetCharactersCharacterIdChatChannelsMuted

/* muted object */
//easyjson:json
type GetCharactersCharacterIdChatChannelsMuted struct {
	AccessorId   int32     `json:"accessor_id,omitempty"`   /* ID of a muted channel member */
	AccessorType string    `json:"accessor_type,omitempty"` /* accessor_type string */
	EndAt        time.Time `json:"end_at,omitempty"`        /* Time at which this accessor will no longer be muted */
	Reason       string    `json:"reason,omitempty"`        /* Reason this accessor is muted */
}
