
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export default function Page() {
	return (
		<main className={`flex min-h-screen flex-col items-center py-6 ${inter.className}`} >
			<h1>admin panel</h1>
		</ main>
	);
}
