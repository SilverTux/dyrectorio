import { SingleFormLayout } from '@app/components/layout'
import DyoButton from '@app/elements/dyo-button'
import { ROUTE_INDEX } from '@app/routes'
import { redirectTo } from '@app/utils'
import { SelfServiceError } from '@ory/kratos-client'
import kratos from '@server/kratos'
import { NextPageContext } from 'next'
import useTranslation from 'next-translate/useTranslation'
import { useRouter } from 'next/dist/client/router'

interface Error {
  message?: string
}

const ErrorPage = (props: SelfServiceError) => {
  const { error: propsError } = props

  const { t } = useTranslation('errors')
  const router = useRouter()

  const error = propsError && 'message' in propsError ? (propsError as Error) : null

  return (
    <SingleFormLayout title={t('oops')}>
      <div className="text-center my-20">
        <h2 className="text-4xl font-extrabold text-blue">{t('oops')}</h2>

        <div className="my-4 text-center text-purple-lightText font-semibold">{error?.message}</div>

        <DyoButton onClick={() => router.back()}>{t('common:goBack')}</DyoButton>

        <div className="mt-16" />
      </div>
    </SingleFormLayout>
  )
}

export default ErrorPage

export async function getServerSideProps(context: NextPageContext) {
  const id = context.query.id as string
  if (!id) {
    return redirectTo(ROUTE_INDEX)
  }

  const error = await kratos.getSelfServiceError(id)

  return {
    props: error.data,
  }
}
