"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language: https://en.wikipedia.org/wiki/Guido_van_Rossum
"""

EXPECTED_BAKE_TIME = 40

def bake_time_remaining(howManyMinutesCookAllready):
    
    """_function to calculate remaining time_

    Returns:
        _type_: _description_
    """
    
    return EXPECTED_BAKE_TIME-howManyMinutesCookAllready

def preparation_time_in_minutes(numberOfLayers):
    """_just simple shit_

    Args:
        numberOfLayers (_type_): _description_

    Returns:
        _type_: _description_
    """
    return numberOfLayers*2



def elapsed_time_in_minutes(numberOfLayers, elapsedBakeTime):
    """_summary_

    Args:
        numberOfLayers (_type_): _description_
        elapsedBakeTime (_type_): _description_

    Returns:
        _type_: _description_
    """

    return preparation_time_in_minutes(numberOfLayers) + elapsedBakeTime