
	$(document).ready(function() {

		window.make_cluster_connections = function() {

		$('div#gtm1').connections({ to: 'div#gtm_slave1', tag: 'to_gtm_slave' });

		$('div#c1').connections({ to: 'div#gp1', tag: 'coord_to_gtm_proxy' });
		$('div#c2').connections({ to: 'div#gp4', tag: 'coord_to_gtm_proxy' });

		$('div#c3').connections({ to: 'div#gp5', tag: 'coord_to_gtm_proxy' });
		$('div#c4').connections({ to: 'div#gp7', tag: 'coord_to_gtm_proxy' });

		$('div#gp1').connections({ to: 'div#gtm1', tag: 'to_gtm' });
		$('div#gp2').connections({ to: 'div#gtm1', tag: 'to_gtm' });
		$('div#gp3').connections({ to: 'div#gtm1', tag: 'to_gtm' });
		$('div#gp4').connections({ to: 'div#gtm1', tag: 'to_gtm' });

		$('div#gp5').connections({ to: 'div#gtm1', tag: 'to_gtm' });
		$('div#gp6').connections({ to: 'div#gtm1', tag: 'to_gtm' });
		$('div#gp7').connections({ to: 'div#gtm1', tag: 'to_gtm' });
		$('div#gp8').connections({ to: 'div#gtm1', tag: 'to_gtm' });

		$('div#dn1').connections({ to: 'div#gp2', tag: 'dn_to_gtm_proxy' });
		$('div#dn2').connections({ to: 'div#gp2', tag: 'dn_to_gtm_proxy' });
		$('div#dn3').connections({ to: 'div#gp3', tag: 'dn_to_gtm_proxy' });
		$('div#dn4').connections({ to: 'div#gp3', tag: 'dn_to_gtm_proxy' });
		$('div#dn5').connections({ to: 'div#gp6', tag: 'dn_to_gtm_proxy' });
		$('div#dn6').connections({ to: 'div#gp6', tag: 'dn_to_gtm_proxy' });
		$('div#dn7').connections({ to: 'div#gp8', tag: 'dn_to_gtm_proxy' });

		$('div#dn1').connections({ to: 'div#ds1', tag: 'dn_to_dn_slave' });
		$('div#dn2').connections({ to: 'div#ds2', tag: 'dn_to_dn_slave' });
		$('div#dn3').connections({ to: 'div#ds3', tag: 'dn_to_dn_slave' });
		$('div#dn4').connections({ to: 'div#ds4', tag: 'dn_to_dn_slave' });
		$('div#dn5').connections({ to: 'div#ds5', tag: 'dn_to_dn_slave' });
		$('div#dn6').connections({ to: 'div#ds6', tag: 'dn_to_dn_slave' });
		$('div#dn7').connections({ to: 'div#ds7', tag: 'dn_to_dn_slave' });
		}

	});
