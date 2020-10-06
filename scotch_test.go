package scotch_test

import (
	"testing"

	"github.com/sawadashota/scotch"
)

func TestScope_Satisfy(t *testing.T) {
	type args struct {
		given scotch.Scope
	}
	cases := map[string]struct {
		required string
		args     args
		want     bool
	}{
		"Perfect match": {
			required: "project:profile=read",
			args: args{
				given: "project:profile=read",
			},
			want: true,
		},
		"Operation contain": {
			required: "project:profile=list",
			args: args{
				given: "project:profile=read,list",
			},
			want: true,
		},
		"Different resource": {
			required: "project:foo=list",
			args: args{
				given: "project:profile=read,list",
			},
			want: false,
		},
		"Not contain": {
			required: "project:profile=write",
			args: args{
				given: "project:profile=read,describe",
			},
			want: false,
		},
		"Resource match": {
			required: "project:profile=read",
			args: args{
				given: "project:profile",
			},
			want: true,
		},
		"Resource match and allow all operation": {
			required: "project:profile=read",
			args: args{
				given: "project:profile=*",
			},
			want: true,
		},
		"Resource prefix match": {
			required: "project:profile=read",
			args: args{
				given: "project:*",
			},
			want: true,
		},
		"Resource hierarchy not match": {
			required: "project:profile=read",
			args: args{
				given: "project:profile:foo=read",
			},
			want: false,
		},
		"Resource hierarchy not match 2": {
			required: "project:profile:foo=read",
			args: args{
				given: "project:profile=read",
			},
			want: false,
		},
		"Resource middle match": {
			required: "project:profile:foo=read",
			args: args{
				given: "project:*:foo",
			},
			want: true,
		},
		"Resource middle match but not match lasting": {
			required: "project:profile:foo=read",
			args: args{
				given: "project:*:bar",
			},
			want: false,
		},
		"Resource part prefix match": {
			required: "project:my-profile=read",
			args: args{
				given: "project:my-*",
			},
			want: true,
		},
		"Resource part prefix match with match operation": {
			required: "project:my-profile=read,write",
			args: args{
				given: "project:my-*=read",
			},
			want: true,
		},
		"Resource part prefix match but not match operation": {
			required: "project:my-profile=read,write",
			args: args{
				given: "project:my-*=list",
			},
			want: false,
		},
		"One all symbol": {
			required: "project:my-profile=read,write",
			args: args{
				given: "*",
			},
			want: true,
		},
		"All symbol for all resource": {
			required: "project:my-profile=read,write",
			args: args{
				given: "*:*",
			},
			want: true,
		},
		"All symbol for all resource with operations": {
			required: "project:my-profile=read,write",
			args: args{
				given: "*:*=*",
			},
			want: true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			required := scotch.New(c.required)

			if got := required.Satisfy(c.args.given); got != c.want {
				t.Errorf("Match() = %v, want %v", got, c.want)
			}
		})
	}
}
