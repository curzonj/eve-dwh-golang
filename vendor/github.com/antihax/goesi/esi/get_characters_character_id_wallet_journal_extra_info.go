/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.5.5
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

/* A list of GetCharactersCharacterIdWalletJournalExtraInfo. */
//easyjson:json
type GetCharactersCharacterIdWalletJournalExtraInfoList []GetCharactersCharacterIdWalletJournalExtraInfo

/* Extra information for different type of transaction */
//easyjson:json
type GetCharactersCharacterIdWalletJournalExtraInfo struct {
	AllianceId          int32  `json:"alliance_id,omitempty"`            /* alliance_id integer */
	CharacterId         int32  `json:"character_id,omitempty"`           /* character_id integer */
	ContractId          int32  `json:"contract_id,omitempty"`            /* contract_id integer */
	CorporationId       int32  `json:"corporation_id,omitempty"`         /* corporation_id integer */
	DestroyedShipTypeId int32  `json:"destroyed_ship_type_id,omitempty"` /* destroyed_ship_type_id integer */
	JobId               int32  `json:"job_id,omitempty"`                 /* job_id integer */
	LocationId          int64  `json:"location_id,omitempty"`            /* location_id integer */
	NpcId               int32  `json:"npc_id,omitempty"`                 /* npc_id integer */
	NpcName             string `json:"npc_name,omitempty"`               /* npc_name string */
	PlanetId            int32  `json:"planet_id,omitempty"`              /* planet_id integer */
	SystemId            int32  `json:"system_id,omitempty"`              /* system_id integer */
	TransactionId       int64  `json:"transaction_id,omitempty"`         /* transaction_id integer */
}
