<div class="content">
  <form method="post" action="/result">
      <div class="form-group">
        <label>Search Place</label>
        <input type="text" class="form-control"
        placeholder="Type place name" name="keyword" value="{{.keyword}}"/>
      </div>
      <button type="submit" class="btn btn-info btn-fill pull-right">Submit</button>
  </form>
  <ul>
    {{range $key, $val := .restaurant}}
        <li>
          <b><span>{{$val.Restaurant.Name}}</span></b>
          <br/>
                    <span>{{$val.Restaurant.Location.Address}}</span>
          <p>
            <button type="submit" class="btn btn-info btn-fill">Add to  omelette</button>
          </p>
        </li>
    {{end}}
    </ul>
</div>
