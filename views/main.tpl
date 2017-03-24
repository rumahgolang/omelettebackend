<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<link rel="icon" type="image/png" href="/static/img/favicon.ico">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />

	<title>Light Bootstrap Dashboard by Creative Tim</title>

	<meta content='width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0' name='viewport' />
    <meta name="viewport" content="width=device-width" />


    <!-- Bootstrap core CSS     -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet" />

    <!-- Animation library for notifications   -->
    <link href="/static/css/animate.min.css" rel="stylesheet"/>

    <!--  Light Bootstrap Table core CSS    -->
    <link href="/static/css/light-bootstrap-dashboard.css" rel="stylesheet"/>


    <!--  CSS for Demo Purpose, don't include it in your project     -->
    <link href="/static/css/demo.css" rel="stylesheet" />


    <!--     Fonts and icons     -->
    <link href="http://maxcdn.bootstrapcdn.com/font-awesome/4.2.0/css/font-awesome.min.css" rel="stylesheet">
    <link href='http://fonts.googleapis.com/css?family=Roboto:400,700,300' rel='stylesheet' type='text/css'>
    <link href="/static/css/pe-icon-7-stroke.css" rel="stylesheet" />

</head>
<body>

<div class="wrapper">
    <div class="sidebar" data-color="purple" data-image="/static/img/sidebar-5.jpg">

    <!--

        Tip 1: you can change the color of the sidebar using: data-color="blue | azure | green | orange | red | purple"
        Tip 2: you can also add an image using data-image tag

    -->

    	<div class="sidebar-wrapper">
            <div class="logo">
                <a class="simple-text">
                    Omelette CMS
                </a>
            </div>

            <ul class="nav">
                <li class="{{.dashboard}}">
                    <a href="/dashboard">
                        <i class="pe-7s-graph"></i>
                        <p>Dashboard</p>
                    </a>
                </li>
                <li class="{{.maps}}">
                    <a href="/places">
                        <i class="pe-7s-map-marker"></i>
                        <p>Places</p>
                    </a>
                </li>
                <li class="{{.merchant}}">
                    <a href="/merchant">
                        <i class="pe-7s-note2"></i>
                        <p>Merchant</p>
                    </a>
                </li>
                <li class="{{.category}}">
                    <a href="/category">
                        <i class="pe-7s-albums"></i>
                        <p>Category Menu</p>
                    </a>
                </li>
                <li class="{{.menu}}">
                    <a href="/menu">
                        <i class="pe-7s-menu"></i>
                        <p>Menus</p>
                    </a>
                </li>
            </ul>
    	</div>
    </div>

    <div class="main-panel">
        <nav class="navbar navbar-default navbar-fixed">
            <div class="container-fluid">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navigation-example-2">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <a class="navbar-brand" href="#">{{.titlePage}}</a>
                </div>
                <div class="collapse navbar-collapse">

                </div>
            </div>
        </nav>

        {{.LayoutContent}}

        <footer class="footer">
            <div class="container-fluid">

                <p class="copyright pull-right">
                    &copy; <script>document.write(new Date().getFullYear())</script> <a href="http://www.creative-tim.com">Creative Tim</a>, made with love for a better web
                </p>
            </div>
        </footer>

    </div>
</div>


</body>

    <!--   Core JS Files   -->
    <script src="/static/js/jquery-1.10.2.js" type="text/javascript"></script>
	<script src="/static/js/bootstrap.min.js" type="text/javascript"></script>

	<!--  Checkbox, Radio & Switch Plugins -->
	<script src="/static/js/bootstrap-checkbox-radio-switch.js"></script>

	<!--  Charts Plugin -->
	<script src="/static/js/chartist.min.js"></script>

    <!--  Notifications Plugin    -->
    <script src="/static/js/bootstrap-notify.js"></script>

    <!-- Light Bootstrap Table Core javascript and methods for Demo purpose -->
	<script src="/static/js/light-bootstrap-dashboard.js"></script>

{{if .isSaved}}
	<script type="text/javascript">
	$(document).ready(function(){
		$.notify({
				icon: 'pe-7s-gift',
				message: "Data has been saved."

			},{
					type: 'info',
					timer: 500
			});

	});
	</script>
	{{end}}
</html>
