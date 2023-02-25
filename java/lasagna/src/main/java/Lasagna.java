public class Lasagna {

    static int STANDARD_TIME_IN_OVEN = 40;
    static int TIME_FOR_PREPARE_ONE_LAYER_OF_LASANGA = 2;

    public int expectedMinutesInOven() {

        return STANDARD_TIME_IN_OVEN;
    }

    public int remainingMinutesInOven(int howManyAlreadyCooked) {

        return STANDARD_TIME_IN_OVEN - howManyAlreadyCooked;
    }

    public int preparationTimeInMinutes(int howManyLayers) {

        return howManyLayers * TIME_FOR_PREPARE_ONE_LAYER_OF_LASANGA;
    }

    public int totalTimeInMinutes(int numberOfLayers, int alreadyCooked) {

        return preparationTimeInMinutes(numberOfLayers) + alreadyCooked;
    }
}
