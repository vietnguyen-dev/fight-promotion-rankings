import { Inter } from "next/font/google";
import { GetServerSideProps, Link } from 'next';

import { iWeightClass } from '../interfaces/weight-classes'
import { iPromotions } from '../interfaces/promotion'
import { iRankings } from '../interfaces/ranking'

const inter = Inter({ subsets: ["latin"] });

interface iPageProps {
	weightClasses: iWeightClass[],
	promotion: iPromotions,
	rankings: iRankings[]
}

export default function Page({ weightClasses, promotion, rankings }: iPageProps) {
	const weightFiltered = weightClasses.filter((weight) => weight.wei_code.includes("CAT") != true)
	return (
		<main
			className={`flex min-h-screen flex-col items-center ${inter.className}`}
		>
			<div className="flex justify-center shadow-bottom w-full mb-6 p-2">
				<img src={promotion.s3_url} alt={`${promotion.name} logo`} width={50} height={50} />
			</div>
			<h1 className='font-extrabold text-2xl text-center px-6'>{promotion.name.toUpperCase()}</h1>
			<h2 className='font-extrabold text-xl px-6 pb-6'>ATHLETE RANKINGS</h2>
			<section className='join join-vertical w-full'>
				{weightFiltered.map((weight) => {
					const ranked = rankings.filter((rank) => rank.weight_class_id === weight.id)
					const champion = ranked[0]
					const rest = ranked.slice(1, ranked.length)
					console.log(rest)
					return (
						<div key={weight.id} className="collapse collapse-arrow join-item border border-base-300 px-3">
							<input type="radio" name="my-accordion-4" />
							<div className="collapse-title">
								<h4 className='text-red-600 text-sm font-bold'>{weight.name.toUpperCase()}</h4>
								{ranked.length > 0 ?
									<>
										<h3 className='font-bold text-xl mb-2'>{champion.athlete_name.toUpperCase()}</h3>
										<h5 className='text-sm'>CHAMPION</h5>
									</> : <h3>Vacant Title</h3>}
							</div>
							<div className="collapse-content">
								{rest.length > 0 ?
									<>
										{rest.map((athlete) =>
											<div key={athlete.id} className='flex items-center'>
												<p className="font-bold text-zinc-600 text-xs mr-3">{athlete.ranking}</p>
												<p className='font-bold text-lg'>{athlete.athlete_name}</p>
											</div>
										)}</> : <p>To Be Determined ...</p>}
							</div>
						</div>
					);
				})}
			</section>
			<footer className="footer footer-center p-10 bg-base-200 text-base-content rounded">
				<nav className="grid grid-flow-col gap-4">
					<a className="link link-hover">About us</a>
					<a className="link link-hover">Contact</a>
					<a className="link link-hover">Jobs</a>
					<a className="link link-hover">Press kit</a>
				</nav>
				<aside>
					<p>Copyright © 2024 - All right reserved by ACME Industries Ltd</p>
				</aside>
			</footer>
		</main>
	);
}

export const getServerSideProps: GetServerSideProps = async (context) => {
	const { query } = context
	const { id } = query

	const weightClassData = await fetch(`http://localhost:8080/weight-classes/${id}`)
	const weightClasses = await weightClassData.json()
	const promotionData = await fetch(`http://localhost:8080/promotions/${id}`)
	const promotion = await promotionData.json()
	const rankingsData = await fetch(`http://localhost:8080/ranking/${id}`)
	const rankings = await rankingsData.json()

	return {
		props: {
			weightClasses: weightClasses,
			promotion: promotion[0],
			rankings: rankings
		}
	}
}
