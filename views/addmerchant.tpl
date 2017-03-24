<div class="content">
            <div class="container-fluid">
                <div class="row">
                    <div class="col-md-6">
                        <div class="card">
                            <div class="header">
                                <h4 class="title">Merchant Profile</h4>
                            </div>
                            <div class="content">
                                <form action="/merchant/{{.profile.ID}}/save" method="post">
                                    <div class="row">
                                        <div class="col-md-12">
                                            <div class="form-group">
                                                <label>Merchant Name</label>
                                                <input type="text" class="form-control" name="merchantname" placeholder="Company" value="{{.profile.Name}}">
                                            </div>
                                        </div>
                                    </div>

                                    <div class="row">
                                        <div class="col-md-6">
                                            <div class="form-group">
                                                <label>Open Close Info</label>
                                                <input type="text" class="form-control" name="opencloseinfo" placeholder="Open Close info" >
                                            </div>
                                        </div>
                                        <div class="col-md-6">
                                            <div class="form-group">
                                                <label>Range Price</label>
                                                <input type="text" class="form-control" name="rangeprice" placeholder="Range Price">
                                            </div>
                                        </div>
                                    </div>

                                    <div class="row">
                                        <div class="col-md-12">
                                            <div class="form-group">
                                                <label>Address</label>
                                                <input type="text" class="form-control" name="address" placeholder="Business Address" value="{{.profile.Location.Address}}">
                                            </div>
                                        </div>
                                    </div>
                                    <input type="text" hidden="hidden" value="{{.profile.Location.Latitude}}" name="latitude" />
                                    <input type="text" hidden="hidden" value="{{.profile.Location.Longitude}}" name="longitude" />
                                    <button type="submit" class="btn btn-info btn-fill pull-right">Update Profile</button>
                                    <div class="clearfix"></div>
                                </form>
                            </div>
                        </div>
                    </div>
                    <div class="col-md-6">
                        <div class="card card-user">
                            <div class="text-center" style="padding:16px">
                                <a href="{{.profile.Url}}" target="_blank">Show more info on Zomato</a>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
        </div>
