<div class="content">
  <div class="container-fluid">
      <div class="row">
          <div class="col-md-8">
              <div class="card">
                  <div class="header">
                      <h4 class="title">Edit Categories</h4>
                  </div>
                  <div class="content">
                      <form action="/category/{{.merchant.ID}}/add" method="post">

                          <div class="row">
                              <div class="col-md-12">
                                  <div class="form-group">
                                      <label>Category Name</label>
                                      <input type="text" class="form-control" name="category_name" placeholder="Category Name">
                                  </div>
                              </div>
                          </div>

                          <button type="submit" class="btn btn-info btn-fill pull-right">Update Category</button>
                          <div class="clearfix"></div>
                      </form>
                  </div>
              </div>
          </div>
          <div class="col-md-4">
                                  <div class="card card-user">
                                      <div class="image">
                                          <img src="https://ununsplash.imgix.net/photo-1431578500526-4d9613015464?fit=crop&fm=jpg&h=300&q=75&w=400" alt="..."/>
                                      </div>
                                      <div class="content">
                                            <h4 class="title">Categories Menu</h4>
                                            <div class="table-responsive">
                                                <table class="table table-hover table-striped">
                                                    <thead>
                                                      <th>ID</th>
                                                      <th>Name</th>
                                                      <th>Action</th>
                                                    </thead>
                                                    <tbody>
                                                      {{range $key, $val := .listCategories.Data}}
                                                        <tr>
                                                          <td>{{$val.ID}}</td>
                                                          <td>{{$val.Name}}</td>
                                                          <td>
                                                            <a href="/category/{{$val.ID}}/edit">Edit</a></td>
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
</div>
</div>
