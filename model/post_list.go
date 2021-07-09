// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
	"sort"
)

type PostList struct {
	Order      []string         `json:"order"`
	Posts      map[string]*Post `json:"posts"`
	NextPostID string           `json:"next_post_id"`
	PrevPostID string           `json:"prev_post_id"`
}

func NewPostList() *PostList {
	return &PostList{
		Order:      make([]string, 0),
		Posts:      make(map[string]*Post),
		NextPostID: "",
		PrevPostID: "",
	}
}

func (o *PostList) ToSlice() []*Post {
	var posts []*Post

	if l := len(o.Posts); l > 0 {
		posts = make([]*Post, 0, l)
	}

	for _, id := range o.Order {
		posts = append(posts, o.Posts[id])
	}
	return posts
}

func (o *PostList) WithRewrittenImageURLs(f func(string) string) *PostList {
	copy := *o
	copy.Posts = make(map[string]*Post)
	for id, post := range o.Posts {
		copy.Posts[id] = post.WithRewrittenImageURLs(f)
	}
	return &copy
}

func (o *PostList) StripActionIntegrations() {
	posts := o.Posts
	o.Posts = make(map[string]*Post)
	for id, post := range posts {
		pcopy := post.Clone()
		pcopy.StripActionIntegrations()
		o.Posts[id] = pcopy
	}
}

func (o *PostList) ToJSON() string {
	copy := *o
	copy.StripActionIntegrations()
	b, err := json.Marshal(&copy)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o *PostList) MakeNonNil() {
	if o.Order == nil {
		o.Order = make([]string, 0)
	}

	if o.Posts == nil {
		o.Posts = make(map[string]*Post)
	}

	for _, v := range o.Posts {
		v.MakeNonNil()
	}
}

func (o *PostList) AddOrder(id string) {

	if o.Order == nil {
		o.Order = make([]string, 0, 128)
	}

	o.Order = append(o.Order, id)
}

func (o *PostList) AddPost(post *Post) {

	if o.Posts == nil {
		o.Posts = make(map[string]*Post)
	}

	o.Posts[post.ID] = post
}

func (o *PostList) UniqueOrder() {
	keys := make(map[string]bool)
	order := []string{}
	for _, postID := range o.Order {
		if _, value := keys[postID]; !value {
			keys[postID] = true
			order = append(order, postID)
		}
	}

	o.Order = order
}

func (o *PostList) Extend(other *PostList) {
	for postID := range other.Posts {
		o.AddPost(other.Posts[postID])
	}

	for _, postID := range other.Order {
		o.AddOrder(postID)
	}

	o.UniqueOrder()
}

func (o *PostList) SortByCreateAt() {
	sort.Slice(o.Order, func(i, j int) bool {
		return o.Posts[o.Order[i]].CreateAt > o.Posts[o.Order[j]].CreateAt
	})
}

func (o *PostList) Etag() string {

	id := "0"
	var t int64 = 0

	for _, v := range o.Posts {
		if v.UpdateAt > t {
			t = v.UpdateAt
			id = v.ID
		} else if v.UpdateAt == t && v.ID > id {
			t = v.UpdateAt
			id = v.ID
		}
	}

	orderID := ""
	if len(o.Order) > 0 {
		orderID = o.Order[0]
	}

	return Etag(orderID, id, t)
}

func (o *PostList) IsChannelID(channelID string) bool {
	for _, v := range o.Posts {
		if v.ChannelID != channelID {
			return false
		}
	}

	return true
}

func PostListFromJSON(data io.Reader) *PostList {
	var o *PostList
	json.NewDecoder(data).Decode(&o)
	return o
}
