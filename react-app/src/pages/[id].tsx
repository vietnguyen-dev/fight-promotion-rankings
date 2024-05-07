import { Inter } from "next/font/google";
import { GetServerSideProps } from 'next';
import { iWeightClass } from '../interfaces/weight-classes'
import { iPromotions } from '../interfaces/promotion'
import { iRankings } from '../interfaces/ranking'

import Accordian from '../components/accordian'

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
			className={`flex min-h-screen flex-col items-center md:w-3/4 md:mx-auto ${inter.className}`}
		>
			<div className="flex justify-center shadow-bottom w-full mb-6 p-2 md:mt-6 md:shadow-xl">
				<img src={promotion.s3_url} alt={`${promotion.name} logo`} width={50} height={50} />
			</div>
			<div className="text-center md:py-12">
				<h1 className="font-extrabold text-2xl px-6">{promotion.name.toUpperCase()}</h1>
				<h2 className="font-extrabold text-xl px-6 pb-6">ATHLETE RANKINGS</h2>
			</div>
			<section className="grid grid-cols-1 w-full md:grid md:grid-cols-4 md:w-2/3 md:mx-auto">
				<Accordian weightClasses={weightFiltered} rankings={rankings} />
			</section>
			<footer className="footer footer-center p-10 bg-base-200 text-base-content rounded">
				<nav className="grid grid-flow-col gap-4">
					<a className="link link-hover">Get Your Own Rankings</a>
					<a className="link link-hover">Login</a>
				</nav>
				<aside>
					<p>Copyright © 2024 - All right reserved by Keith Nguyen 2024</p>
				</aside>
			</footer>
		</main >
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
