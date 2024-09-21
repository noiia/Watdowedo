package basetemplate

import "html/template"

type Page struct {
	Header template.HTML
	Navbar template.HTML
	Footer template.HTML
}

func LoadBase() Page {
	return Page{
		Header: `<meta charset="UTF-8" />
			<meta http-equiv="X-UA-Compatible" content="IE=edge" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
			<link rel="stylesheet" type="text/css" href="../static/css/output.css"/>`,

		Navbar: `
		<div id="navbar-menu" class="hidden">
				<div class="grid grid-cols-2 gap-auto">
					<a href="./home">
						<img
							src="../static/images/logos/watdowedo_full_logo.png"
							alt="logo_watdowedo"
						/>
					</a>
					<div
						class="grid grid-cols-[150px_150px_150px] xl:gap-10 list-none justify-self-end mr-5"
					>
						<a
							href=""
							class="m-auto box-border h-[75px] w-[150px] rounded hover:bg-slate-300 transition-all flex justify-center items-center"
						>
							<li class="m-auto">Trips</li>
						</a>
						<a
							href=""
							class="m-auto box-border h-[75px] w-[150px] rounded hover:bg-slate-300 transition-all flex justify-center items-center"
						>
							<li class="m-auto">Trip Builder</li>
						</a>
						<a
							href=""
							class="m-auto box-border h-[75px] w-[150px] rounded bg-primary hover:scale-110 transition-all flex justify-center items-center"
						>
							<li class="m-auto">Sign in</li>
						</a>
					</div>
				</div>
			</div>
			<div id="navbar-burger" class="hidden">
				<div class="box-border h-[50px] w-[50px]">
					<img src="../static/images/icon/menu.png" alt="logo_watdowedo" />
				</div>			
			</div>
			<script
				type="text/javascript"
				src="../static/js/basetemplate.js"
			></script>
			`,
		Footer: `
		<div class="bg-neutral-content bottom-0 inset-x-0 pt-6 pb-8">
			<div class="grid grid-cols-2 over-cellphone:grid-cols-4 py-4 gap-4">
					<div>
						<img
							src="../static/images/logos/watdowedo_tree_logo.png"
							alt="logo de l'entreprise"
							class="mx-auto"
						/>
					</div>
					<div>
						<h1 class="text-center font-semibold text-xl pb-6">Société</h1>
						<ul class="text-center">
							<li>Condition d'utilisation</li>
							<li>Politique de confidentialité</li>
							<li>Politique de cookies</li>
							<li>A propos de nous</li>
						</ul>
					</div>
					<div>
						<h1 class="text-center font-semibold text-xl pb-6">Aide</h1>
						<ul class="text-center">
							<li>FAQs</li>
							<li>Nous contacter</li>
						</ul>
					</div>
					<div>
						<h1 class="text-center font-semibold text-xl pb-6">Follow us</h1>
						<ul class="grid grid-cols-1 sm:grid-cols-2 gap-4 p-4">
							<a href="https://www.facebook.com">
								<li
									class="mx-auto box-border h-auto max-h-14 over-cellphone:h-14 w-auto max-w-14 over-cellphone:w-14"
								>
									<img
										src="../static/images/logos/logo_facebook.png"
										alt="logo facebook"
									/>
								</li>
							</a>
							<a href="https://www.instagram.com">
								<li
									class="mx-auto box-border h-auto max-h-14 over-cellphone:h-14 w-auto max-w-14 over-cellphone:w-14"
								>
									<img
										src="../static/images/logos/logo_insta.png"
										alt="logo instagram"
									/>
								</li>
							</a>
							<a href="https://www.youtube.com">
								<li
									class="mx-auto box-border h-auto max-h-14 over-cellphone:h-14 w-auto max-w-14 over-cellphone:w-14"
								>
									<img
										src="../static/images/logos/logo_ytb.png"
										alt="logo youtube"
									/>
								</li>
							</a>
							<a href="https://www.twitter.com">
								<li
									class="mx-auto box-border h-auto max-h-14 over-cellphone:h-14 w-auto max-w-14 over-cellphone:w-14"
								>
									<img src="../static/images/logos/logo_X.png" alt="logo X" />
								</li>
							</a>
						</ul>
					</div>
				</div>
				<a href="https://www.flaticon.com/free-icons/walk" title="walk icons"
					>Icons created by Freepik - Flaticon</a
				>
				<a href="https://www.flaticon.com/fr/icones-gratuites/angleterre" title="angleterre icônes">Angleterre icônes créées par IconMarketPK - Flaticon</a>
			</div>
		`,
	}
}
