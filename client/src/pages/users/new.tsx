import { Button, TextField } from '@material-ui/core'
import Router from 'next/router'
import { useState } from 'react'
import { FormProvider, useForm } from 'react-hook-form'

import formStyles from '../../styles/form.module.scss'

const New = () => {
  const methods = useForm()
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = methods
  const [errorMessage, setErrorMessage] = useState('')

  const onSubmit = handleSubmit(async ({ name }) => {
    if (errorMessage) setErrorMessage('')

    console.info(name)

    try {
      Router.push('/crud-example')
    } catch (error) {
      console.error(error)
      setErrorMessage(error.message)
    }
  })

  const { ref: inputRef } = register('name', {
    required: 'Name is required',
  })

  return (
    <>
      <h1>Create New User</h1>

      <FormProvider {...methods}>
        <form onSubmit={onSubmit} className={formStyles.form}>
          <div>
            <TextField type="text" name="name" label="Name" placeholder="Input Your Name" inputRef={inputRef} />
            {errors.name && (
              <span role="alert" className={formStyles.error}>
                {errors.name.message}
              </span>
            )}
          </div>

          <div className={formStyles.submit}>
            <Button type="submit" variant="contained">
              Create
            </Button>
          </div>
        </form>
      </FormProvider>

      {errorMessage && (
        <p role="alert" className={formStyles.errorMessage}>
          {errorMessage}
        </p>
      )}
    </>
  )
}

// eslint-disable-next-line import/no-default-export
export default New
