// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostListJSON(t *testing.T) {

	pl := PostList{}
	p1 := &Post{ID: NewID(), Message: NewID()}
	pl.AddPost(p1)
	p2 := &Post{ID: NewID(), Message: NewID()}
	pl.AddPost(p2)

	pl.AddOrder(p1.ID)
	pl.AddOrder(p2.ID)

	json := pl.ToJSON()
	rpl := PostListFromJSON(strings.NewReader(json))

	assert.Equal(t, p1.Message, rpl.Posts[p1.ID].Message, "failed to serialize p1 message")
	assert.Equal(t, p2.Message, rpl.Posts[p2.ID].Message, "failed to serialize p2 message")
	assert.Equal(t, p2.ID, rpl.Order[1], "failed to serialize p2 Id")
}

func TestPostListExtend(t *testing.T) {
	p1 := &Post{ID: NewID(), Message: NewID()}
	p2 := &Post{ID: NewID(), Message: NewID()}
	p3 := &Post{ID: NewID(), Message: NewID()}

	l1 := PostList{}
	l1.AddPost(p1)
	l1.AddOrder(p1.ID)
	l1.AddPost(p2)
	l1.AddOrder(p2.ID)

	l2 := PostList{}
	l2.AddPost(p3)
	l2.AddOrder(p3.ID)

	l2.Extend(&l1)

	// should not changed l1
	assert.Len(t, l1.Posts, 2)
	assert.Len(t, l1.Order, 2)

	// should extend l2
	assert.Len(t, l2.Posts, 3)
	assert.Len(t, l2.Order, 3)

	// should extend the Order of l2 correctly
	assert.Equal(t, l2.Order[0], p3.ID)
	assert.Equal(t, l2.Order[1], p1.ID)
	assert.Equal(t, l2.Order[2], p2.ID)

	// extend l2 again
	l2.Extend(&l1)
	// extending l2 again should not changed l1
	assert.Len(t, l1.Posts, 2)
	assert.Len(t, l1.Order, 2)

	// extending l2 again should extend l2
	assert.Len(t, l2.Posts, 3)
	assert.Len(t, l2.Order, 3)

	// p3 could be last unread
	p4 := &Post{ID: NewID(), Message: NewID()}
	p5 := &Post{ID: NewID(), RootID: p1.ID, Message: NewID()}
	p6 := &Post{ID: NewID(), RootID: p1.ID, Message: NewID()}

	// Create before and after post list where p3 could be last unread

	// Order has 2 but Posts are 4 which includes additional 2 comments (p5 & p6) to parent post (p1)
	beforePostList := PostList{
		Order: []string{p1.ID, p2.ID},
		Posts: map[string]*Post{p1.ID: p1, p2.ID: p2, p5.ID: p5, p6.ID: p6},
	}

	// Order has 3 but Posts are 4 which includes 1 parent post (p1) of comments (p5 & p6)
	afterPostList := PostList{
		Order: []string{p4.ID, p5.ID, p6.ID},
		Posts: map[string]*Post{p1.ID: p1, p4.ID: p4, p5.ID: p5, p6.ID: p6},
	}

	beforePostList.Extend(&afterPostList)

	// should not changed afterPostList
	assert.Len(t, afterPostList.Posts, 4)
	assert.Len(t, afterPostList.Order, 3)

	// should extend beforePostList
	assert.Len(t, beforePostList.Posts, 5)
	assert.Len(t, beforePostList.Order, 5)

	// should extend the Order of beforePostList correctly
	assert.Equal(t, beforePostList.Order[0], p1.ID)
	assert.Equal(t, beforePostList.Order[1], p2.ID)
	assert.Equal(t, beforePostList.Order[2], p4.ID)
	assert.Equal(t, beforePostList.Order[3], p5.ID)
	assert.Equal(t, beforePostList.Order[4], p6.ID)
}

func TestPostListSortByCreateAt(t *testing.T) {
	pl := PostList{}
	p1 := &Post{ID: NewID(), Message: NewID(), CreateAt: 2}
	pl.AddPost(p1)
	p2 := &Post{ID: NewID(), Message: NewID(), CreateAt: 1}
	pl.AddPost(p2)
	p3 := &Post{ID: NewID(), Message: NewID(), CreateAt: 3}
	pl.AddPost(p3)

	pl.AddOrder(p1.ID)
	pl.AddOrder(p2.ID)
	pl.AddOrder(p3.ID)

	pl.SortByCreateAt()

	assert.EqualValues(t, pl.Order[0], p3.ID)
	assert.EqualValues(t, pl.Order[1], p1.ID)
	assert.EqualValues(t, pl.Order[2], p2.ID)
}

func TestPostListToSlice(t *testing.T) {
	pl := PostList{}
	p1 := &Post{ID: NewID(), Message: NewID(), CreateAt: 2}
	pl.AddPost(p1)
	p2 := &Post{ID: NewID(), Message: NewID(), CreateAt: 1}
	pl.AddPost(p2)
	p3 := &Post{ID: NewID(), Message: NewID(), CreateAt: 3}
	pl.AddPost(p3)

	pl.AddOrder(p1.ID)
	pl.AddOrder(p2.ID)
	pl.AddOrder(p3.ID)

	want := []*Post{p1, p2, p3}

	assert.Equal(t, want, pl.ToSlice())
}
