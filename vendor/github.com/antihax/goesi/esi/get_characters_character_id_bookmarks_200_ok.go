/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.6.0
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

/* A list of GetCharactersCharacterIdBookmarks200Ok. */
//easyjson:json
type GetCharactersCharacterIdBookmarks200OkList []GetCharactersCharacterIdBookmarks200Ok

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdBookmarks200Ok struct {
	BookmarkId int64                                   `json:"bookmark_id,omitempty"` /* bookmark_id integer */
	CreateDate time.Time                               `json:"create_date,omitempty"` /* create_date string */
	CreatorId  int32                                   `json:"creator_id,omitempty"`  /* creator_id integer */
	FolderId   int32                                   `json:"folder_id,omitempty"`   /* folder_id integer */
	Memo       string                                  `json:"memo,omitempty"`        /* memo string */
	Note       string                                  `json:"note,omitempty"`        /* note string */
	OwnerId    int32                                   `json:"owner_id,omitempty"`    /* owner_id integer */
	Target     GetCharactersCharacterIdBookmarksTarget `json:"target,omitempty"`
}
