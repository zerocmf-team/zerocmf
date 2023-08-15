syntax = "v1"

info (
	title: // TODO: add title
	desc: // TODO: add description
	author: "{{.gitUser}}"
	email: "{{.gitEmail}}"
)

typ GetReq {

}

service {{.serviceName}} {
	@handler Get
    get / ()

    @handler Show
    get /:id ()

    @handler Store
    post / ()

    @handler Edit
    post /:id ()

    @handler Delete
    delete /:id ()

    @handler BatchDelete
    delete / ()
}
