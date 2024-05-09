import { SignIn } from "@clerk/nextjs";

export default function Page() {
	return (
		<section className='h-screen flex justify-center mt-6'>
			<SignIn path="/sign-in" forceRedirectUrl="/admin" />
		</section>
	);
}
