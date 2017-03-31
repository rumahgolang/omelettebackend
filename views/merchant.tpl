<div class="content">
            <div class="container-fluid">
                <div class="row">
                    <div class="col-md-12">
                        <div class="card">
                            <div class="header">
                                <h4 class="title">List of Merchants</h4>
                                <p class="category">This is list of registered merchant</p>
                            </div>
                            <div class="content table-responsive table-full-width">
                                <table class="table table-hover table-striped">
                                    <thead>
                                      <th>ID</th>
                                    	<th>Name</th>
                                    	<th>Location</th>
                                    	<th>Action</th>
                                    </thead>
                                    <tbody>
                                      {{range $key, $val := .merchants.Data}}
                                        <tr>
                                        	<td>{{$val.ID}}</td>
                                        	<td>{{$val.Name}}</td>
                                        	<td>{{$val.Location}}</td>
                                        	<td>
                                            Edit<br/>
                                            <a href="/category/{{$val.ID}}">Added Menu</a></td>
                                        </tr>
                                      {{end}}
                                    </tbody>
                                </table>

                            </div>
                        </div>
                    </div>
                  </div>
            </div>
        </div>
