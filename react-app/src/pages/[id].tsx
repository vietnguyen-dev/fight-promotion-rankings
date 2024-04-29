import { Inter } from "next/font/google";
import { GetServerSideProps } from 'next';
import Link from 'next/link'
import { useRouter } from 'next/router'

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
	const router = useRouter()

	const weightFiltered = weightClasses.filter((weight) => weight.wei_code.includes("CAT") != true)
	return (
		<main
			className={`flex min-h-screen flex-col items-center ${inter.className}`}
		>
			<div className="flex justify-center shadow-bottom w-full mb-6 p-2">
				<img src={promotion.s3_url} alt={`${promotion.name} logo`} width={50} height={50} />
			</div>
			<h1 className="font-extrabold text-2xl text-center px-6">{promotion.name.toUpperCase()}</h1>
			<h2 className="font-extrabold text-xl px-6 pb-6">ATHLETE RANKINGS</h2>
			<section className="join join-vertical w-full">
				<Accordian weightClasses={weightFiltered} rankings={rankings} />
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
