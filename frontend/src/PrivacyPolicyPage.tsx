import { PRIVACY_POLICY } from './legal'

export default function PrivacyPolicyPage() {
  return (
    <div className="min-h-screen bg-[#0b1220] text-gray-200">
      <div className="mx-auto max-w-4xl px-6 py-10">
        <h1 className="text-3xl font-semibold">Privacy Policy</h1>
        <div className="mt-6 whitespace-pre-wrap text-sm leading-relaxed text-white/70">{PRIVACY_POLICY}</div>
      </div>
    </div>
  )
}
